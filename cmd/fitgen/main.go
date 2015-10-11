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
)

const (
	fitPkgImportPath = "github.com/tormoder/fit"

	types    = "types"
	msgs     = "messages"
	profilef = "profile"

	workbookNameXLS  = "Profile.xls"
	workbookNameXLSX = "Profile.xlsx"
)

func main() {
	keep := flag.Bool(
		"keep",
		false,
		"don't delete intermediary workbook and csv files from profile directory",
	)
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
		fmt.Fprintf(os.Stderr, "usage: fitgenprofile [-keep] [path to sdk zip, xls or xlsx file]\n")
		flag.PrintDefaults()
	}

	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(2)
	}

	log.SetFlags(0)
	log.SetPrefix("fitgen:\t")

	_, err := exec.LookPath("python3")
	if err != nil {
		log.Fatal("can't find python3 in $PATH, needed for xls to csv conversion")
	}

	fitSrcDir, err := goPackagePath(fitPkgImportPath)
	if err != nil {
		log.Fatal("can't find fit package root src directory")
	}
	log.Println("root src directory:", fitSrcDir)
	profileDir := filepath.Join(fitSrcDir, "profile")

	var (
		isZip         bool
		workbookPath  string
		sdkVersion    string
		stringerInput string

		goTypes map[string]*profile.Type
		goMsgs  []*profile.Msg

		pyScript         = filepath.Join(fitSrcDir, "lib/python/xls_to_csv.py")
		typesCSVOut      = filepath.Join(profileDir, types+".csv")
		msgsCSVOut       = filepath.Join(profileDir, msgs+".csv")
		msgsGoOut        = filepath.Join(fitSrcDir, msgs+".go")
		typesGoOut       = filepath.Join(fitSrcDir, types+".go")
		profileGoOut     = filepath.Join(fitSrcDir, profilef+".go")
		typesStringGoOut = filepath.Join(fitSrcDir, types+"_string.go")
		stringerPath     = filepath.Join(fitSrcDir, "cmd/stringer/stringer.go")
	)

	input := flag.Arg(0)
	inputExt := filepath.Ext(input)
	switch inputExt {
	case ".zip":
		isZip = true
	case ".xls", ".xlsx":
		workbookPath = input
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

	if isZip {
		workbookPath, err = writeWorkbookToFile(input, profileDir)
		if err != nil {
			log.Fatalln("write-xls-to-file:", err)
		}
	}

	err = generateCSV(pyScript, workbookPath, typesCSVOut, msgsCSVOut)
	if err != nil {
		goto fatal
	}

	goTypes, stringerInput, err = generateTypes(sdkVersion, typesCSVOut, typesGoOut)
	if err != nil {
		goto fatal
	}

	goMsgs, err = generateMsgs(sdkVersion, msgsCSVOut, msgsGoOut, goTypes)
	if err != nil {
		goto fatal
	}

	err = generateProfile(sdkVersion, goTypes, goMsgs, profileGoOut, *jmptable)
	if err != nil {
		goto fatal
	}

	err = runStringerOnTypes(stringerPath, fitSrcDir, typesStringGoOut, stringerInput)
	if err != nil {
		goto fatal
	}

	err = runAllTests(fitPkgImportPath)
	if err != nil {
		goto fatal
	}

	err = logMesgNumVsMsgs(goTypes, goMsgs)
	if err != nil {
		goto fatal
	}

	cleanup(workbookPath, typesCSVOut, msgsCSVOut, *keep, isZip)
	log.Println("done")
	os.Exit(0)

fatal:
	cleanup(workbookPath, typesCSVOut, msgsCSVOut, *keep, isZip)
	log.Fatal(err)
}

func parseSDKVersion(zipFilePath string) string {
	// Brittle.
	// TODO: Maybe parse 'c/fit.h' with regexp instead.
	_, file := filepath.Split(zipFilePath)
	ver := strings.TrimSuffix(file, ".zip")
	return strings.TrimPrefix(ver, "FitSDKRelease_")
}

func writeWorkbookToFile(inputPath, profilePath string) (string, error) {
	r, err := zip.OpenReader(inputPath)
	if err != nil {
		return "", fmt.Errorf("error opening sdk zip file: %v", err)
	}
	defer r.Close()

	var workbook *zip.File
	for _, f := range r.File {
		if f.Name == workbookNameXLS {
			workbook = f
			break
		}
		if f.Name == workbookNameXLSX {
			workbook = f
			break
		}
	}
	if workbook == nil {
		return "", fmt.Errorf(
			"no file named %q or %q found in zip archive",
			workbookNameXLS,
			workbookNameXLSX,
		)
	}

	rc, err := workbook.Open()
	defer rc.Close()
	if err != nil {
		return "", err
	}
	b, err := ioutil.ReadAll(rc)
	if err != nil {
		return "", err
	}
	fpath := filepath.Join(profilePath, workbook.Name)
	err = ioutil.WriteFile(fpath, b, 0644)
	if err != nil {
		return "", err
	}
	log.Println("write-xls-to-file: done")
	return fpath, nil
}

func generateCSV(pyScript, workbookPath, typesCSVOut, msgsCSVOut string) error {
	cmdConvertTypes := exec.Command(pyScript, "-wb", workbookPath, "-sh", "Types", "-out", typesCSVOut)
	output, err := cmdConvertTypes.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error converting types from xls to csv: %v\n%s", err, output)
	}
	cmdConvertMsgs := exec.Command(pyScript, "-wb", workbookPath, "-sh", "Messages", "-out", msgsCSVOut)
	output, err = cmdConvertMsgs.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error converting messages from xls to csv: %v\n%s", err, output)
	}
	log.Println("xls-csv-gen: done")
	return nil
}

func generateTypes(sdkVersion, input, output string) (map[string]*profile.Type, string, error) {
	file, err := os.Open(input)
	if err != nil {
		return nil, "", fmt.Errorf("typegen: error opening input file: %v", err)
	}
	defer file.Close()

	parser, err := profile.NewTypeParser(file)
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

	generator := profile.NewGenerator(sdkVersion)
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

func generateMsgs(sdkVersion, input, output string, types map[string]*profile.Type) ([]*profile.Msg, error) {
	file, err := os.Open(input)
	if err != nil {
		return nil, fmt.Errorf("msggen: error opening input file: %v", err)
	}
	defer file.Close()

	parser, err := profile.NewMsgParser(file)
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

	generator := profile.NewGenerator(sdkVersion)
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

func generateProfile(sdkVersion string, types map[string]*profile.Type, msgs []*profile.Msg, output string, jmptable bool) error {
	generator := profile.NewGenerator(sdkVersion)
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
func runStringerOnTypes(stringerPath, fitSrcDir, goTypesStringOut, fitTypes string) error {
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
	testCmd := exec.Command("go", "test", pkgDir+"/...")
	output, err := testCmd.CombinedOutput()
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

func cleanup(workbook, types, msgs string, keep, isZip bool) {
	if keep {
		return
	}
	log.Println("cleaning up")
	os.Remove(types)
	os.Remove(msgs)
	if isZip {
		os.Remove(workbook)
	}
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
