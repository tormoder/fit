package main

import (
	"archive/zip"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"

	"github.com/tormoder/fit/cmd/fitgen/internal/profile"

	"github.com/tealeg/xlsx"
)

const (
	fitPkgImportPath = "github.com/tormoder/fit"

	typesSheetIndex = 0
	msgsSheetIndex  = 1
	msgsNumCells    = 16

	workbookNameXLS  = "Profile.xls"
	workbookNameXLSX = "Profile.xlsx"
)

var (
	fitSrcDir        string
	msgsGoOut        string
	typesGoOut       string
	profileGoOut     string
	typesStringGoOut string
	stringerPath     string
	generator        *profile.Generator
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("fitgen:\t")

	var err error
	fitSrcDir, err = goPackagePath(fitPkgImportPath)
	if err != nil {
		log.Fatalf("can't find fit package root src directory for %q", fitPkgImportPath)
	}
	log.Println("root src directory:", fitSrcDir)

	msgsGoOut = filepath.Join(fitSrcDir, "messages.go")
	typesGoOut = filepath.Join(fitSrcDir, "types.go")
	profileGoOut = filepath.Join(fitSrcDir, "profile.go")
	typesStringGoOut = filepath.Join(fitSrcDir, "types_string.go")
	stringerPath = filepath.Join(fitSrcDir, "cmd/stringer/stringer.go")
}

func main() {
	sdkOverride := flag.String(
		"sdk",
		"",
		"provide or override SDK version printed in generated code",
	)
	jmptable := flag.Bool(
		"jmptable",
		true,
		"use jump tables for profile message and field lookups, otherwise use switches",
	)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: fitgen [flags] [path to sdk zip, xls or xlsx file]\n")
		flag.PrintDefaults()
	}

	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(2)
	}

	var sdkVersion string
	isZip := false
	input := flag.Arg(0)
	inputExt := filepath.Ext(input)
	switch inputExt {
	case ".zip":
		isZip = true
	case ".xls", ".xlsx":
	default:
		log.Fatalln("input file must be of type [.zip | .xls | .xlsx], got:", inputExt)
	}

	switch {
	case *sdkOverride != "":
		sdkVersion = *sdkOverride
	case isZip:
		sdkVersion = parseSDKVersion(input)
	default:
		sdkVersion = "Unknown"
	}
	log.Println("sdk version:", sdkVersion)

	generator = profile.NewGenerator(sdkVersion)

	typeData, msgData, err := parseProfileWorkbook(input)
	if err != nil {
		log.Fatal(err)
	}

	goTypes, stringerInput, err := generateTypes(typeData, typesGoOut)
	if err != nil {
		log.Fatal(err)
	}

	goMsgs, err := generateMsgs(msgData, msgsGoOut, goTypes)
	if err != nil {
		log.Fatal(err)
	}

	err = generateProfile(goTypes, goMsgs, profileGoOut, *jmptable)
	if err != nil {
		log.Fatal(err)
	}

	err = runStringerOnTypes(stringerPath, typesStringGoOut, stringerInput)
	if err != nil {
		log.Fatal(err)
	}

	err = runAllTests(fitPkgImportPath)
	if err != nil {
		log.Fatal(err)
	}

	err = logMesgNumVsMsgs(goTypes, goMsgs)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("done")
}

func parseSDKVersion(zipFilePath string) string {
	// Brittle.
	// TODO: Maybe parse 'c/fit.h' with regexp instead.
	_, file := filepath.Split(zipFilePath)
	ver := strings.TrimSuffix(file, ".zip")
	return strings.TrimPrefix(ver, "FitSDKRelease_")
}

func parseProfileWorkbook(inputPath string) (typeData, msgData [][]string, err error) {
	r, err := zip.OpenReader(inputPath)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening sdk zip file: %v", err)
	}
	defer r.Close()

	var wfile *zip.File
	for _, f := range r.File {
		if f.Name == workbookNameXLS {
			wfile = f
			break
		}
		if f.Name == workbookNameXLSX {
			wfile = f
			break
		}
	}
	if wfile == nil {
		return nil, nil, fmt.Errorf(
			"no file named %q or %q found in zip archive",
			workbookNameXLS, workbookNameXLSX)
	}

	rc, err := wfile.Open()
	defer rc.Close()
	if err != nil {
		return nil, nil, err
	}
	b, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, nil, err
	}

	workbook, err := xlsx.OpenBinary(b)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening profile workbook: %v", err)
	}

	// file.ToSlice from the xlsx library adjusted to ignore formatting errors.
	var output = [][][]string{}
	for _, sheet := range workbook.Sheets {
		s := [][]string{}
		for _, row := range sheet.Rows {
			if row == nil {
				continue
			}
			r := []string{}
			for _, cell := range row.Cells {
				str, err := cell.String()
				if err != nil {
					// The profile message sheet has formatting errors.
					// Ignore those cells and use the raw values.
				}
				r = append(r, str)
			}
			s = append(s, r)
		}
		output = append(output, s)
	}

	typeData = output[typesSheetIndex]
	msgData = output[msgsSheetIndex]

	log.Println("parse workbook: done")
	return typeData, msgData, nil
}

func generateTypes(typeData [][]string, output string) (map[string]*profile.Type, string, error) {
	parser, err := profile.NewTypeParser(typeData)
	if err != nil {
		return nil, "", fmt.Errorf("typegen: error creating parser: %v", err)
	}

	var ptypes []*profile.PType
	for {
		t, perr := parser.ParseType()
		if perr == io.EOF {
			break
		}
		if perr != nil {
			return nil, "", fmt.Errorf("typegen: error parsing types: %v", perr)
		}
		ptypes = append(ptypes, t)
	}

	types, err := profile.TransformTypes(ptypes)
	if err != nil {
		return nil, "", fmt.Errorf("typegen: error transforming types: %v", err)
	}

	source, err := generator.GenerateTypes(types)
	if err != nil {
		return nil, "", fmt.Errorf("typegen: error generating source: %v", err)

	}

	if err = ioutil.WriteFile(output, source, 0644); err != nil {
		return nil, "", fmt.Errorf("typegen: error writing output file: %v", err)
	}

	tkeys := make([]string, 0, len(types))
	for tkey := range types {
		tkeys = append(tkeys, tkey)
	}
	sort.Strings(tkeys)

	var atn bytes.Buffer
	for _, tkey := range tkeys {
		t := types[tkey]
		atn.WriteString(t.CCName)
		atn.WriteByte(',')
	}

	allTypeNames := atn.Bytes()
	allTypeNames = allTypeNames[:len(allTypeNames)-1] // last comma

	log.Println("typegen: success")
	return types, string(allTypeNames), nil
}

func generateMsgs(msgData [][]string, output string, types map[string]*profile.Type) ([]*profile.Msg, error) {
	parser, err := profile.NewMsgParser(msgData)
	if err != nil {
		return nil, fmt.Errorf("msggen: error creating parser: %v", err)
	}

	var pmsgs []*profile.PMsg
	for {
		m, perr := parser.ParseMsg()
		if perr == io.EOF {
			break
		}
		if perr != nil {
			return nil, fmt.Errorf("msggen: error parsing msgs: %v", perr)
		}
		pmsgs = append(pmsgs, m)
	}

	msgs, err := profile.TransformMsgs(pmsgs, types)
	if err != nil {
		return nil, fmt.Errorf("msggen: error transforming msgs: %v", err)
	}

	source, err := generator.GenerateMsgs(msgs)
	if err != nil {
		return nil, fmt.Errorf("msggen: error generating source: %v", err)

	}

	if err = ioutil.WriteFile(output, source, 0644); err != nil {
		return nil, fmt.Errorf("msggen: error writing output file: %v", err)
	}

	log.Println("msggen: success")
	return msgs, nil
}

func generateProfile(types map[string]*profile.Type, msgs []*profile.Msg, output string, jmptable bool) error {
	source, err := generator.GenerateProfile(types, msgs, jmptable)
	if err != nil {
		return fmt.Errorf("profilegen: error generating source: %v", err)

	}

	if err = ioutil.WriteFile(output, source, 0644); err != nil {
		return fmt.Errorf("profilegen: error writing output file: %v", err)
	}

	log.Println("profilegen: success")
	return nil
}

func runStringerOnTypes(stringerPath, goTypesStringOut, fitTypes string) error {
	stringerCmd := exec.Command(
		"go",
		"run",
		stringerPath,
		"-trimprefix",
		"-type", fitTypes,
		"-output",
		goTypesStringOut,
		fitSrcDir,
	)

	output, err := stringerCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("stringer: error running on types: %v\n%s", err, output)
	}

	log.Println("stringer: types done")
	return nil
}

func runAllTests(pkgDir string) error {
	listCmd := exec.Command("go", "list", pkgDir+"/...")
	output, err := listCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("go list: fail: %v\n%s", err, output)
	}

	splitted := strings.Split(string(output), "\n")
	var goTestArgs []string
	// Command
	goTestArgs = append(goTestArgs, "test")
	// Pacakges
	for _, s := range splitted {
		if strings.Contains(s, "/vendor/") {
			continue
		}
		goTestArgs = append(goTestArgs, s)
	}

	testCmd := exec.Command("go", goTestArgs...)
	output, err = testCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("go test: fail: %v\n%s", err, output)
	}

	log.Println("go test: pass")
	return nil
}

func logMesgNumVsMsgs(types map[string]*profile.Type, msgs []*profile.Msg) error {
	mesgNum, found := types["MesgNum"]
	if !found {
		return errors.New("mesgnum-vs-#msgs: can't find MesgNum type")
	}

	nMesgNum := len(mesgNum.Values) - 2 // Skip range min/max
	diff := nMesgNum - len(msgs)

	log.Println("mesgnum-vs-msgs: #mesgnum values:", nMesgNum)
	log.Println("mesgnum-vs-msgs: #generated messages:", len(msgs))

	if diff == 0 {
		return nil
	}

	msgsMap := make(map[string]*profile.Msg)
	for _, msg := range msgs {
		msgsMap[msg.CCName] = msg
	}

	var mdiff []string
	for _, mnv := range mesgNum.Values {
		if strings.HasPrefix(mnv.Name, "MfgRange") {
			continue
		}
		_, ok := msgsMap[mnv.Name]
		if !ok {
			mdiff = append(mdiff, mnv.Name)
		}
	}

	log.Println("mesgnum-vs-msgs: #mesgnum values != #generated messages, diff:", diff)
	log.Println("mesgnum-vs-msgs: remember to verify map in mappings.go for the following message(s):")
	for _, mvn := range mdiff {
		log.Printf("mesgnum-vs-msgs: ----> mesgnum %q has no corresponding message\n", mvn)
	}
	log.Println("mesgnum-vs-msgs: this may be automated in the future")

	return nil
}

func goPackagePath(pkg string) (path string, err error) {
	gp := os.Getenv("GOPATH")
	if gp == "" {
		return path, os.ErrNotExist
	}
	for _, p := range filepath.SplitList(gp) {
		dir := filepath.Join(p, "src", filepath.FromSlash(pkg))
		fi, err := os.Stat(dir)
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			return "", err
		}
		if !fi.IsDir() {
			continue
		}
		return dir, nil
	}
	return path, os.ErrNotExist
}
