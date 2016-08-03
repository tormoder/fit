package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/tormoder/fit/cmd/fitgen/internal/profile"
)

const fitPkgImportPath = "github.com/tormoder/fit"

const (
	workbookNameXLS  = "Profile.xls"
	workbookNameXLSX = "Profile.xlsx"
)

func main() {
	l := log.New(os.Stdout, "fitgen:\t", 0)

	fitSrcDir, err := goPackagePath(fitPkgImportPath)
	if err != nil {
		l.Fatalf("can't find fit package root src directory for %q", fitPkgImportPath)
	}
	l.Println("root src directory:", fitSrcDir)

	var (
		messagesOut    = filepath.Join(fitSrcDir, "messages.go")
		typesOut       = filepath.Join(fitSrcDir, "types.go")
		profileOut     = filepath.Join(fitSrcDir, "profile.go")
		stringerPath   = filepath.Join(fitSrcDir, "cmd/stringer/stringer.go")
		typesStringOut = filepath.Join(fitSrcDir, "types_string.go")
	)

	sdkOverride := flag.String(
		"sdk",
		"",
		"provide or override SDK version printed in generated code",
	)
	switches := flag.Bool(
		"jmptable",
		false,
		"use switches instead jump tables for profile message and field lookups",
	)
	timestamp := flag.Bool(
		"timestamp",
		true,
		"add generation timestamp to generated code",
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

	var (
		inputData []byte
		input     = flag.Arg(0)
		inputExt  = filepath.Ext(input)
	)

	switch inputExt {
	case ".zip":
		inputData, err = readDataFromZIP(input)
	case ".xls", ".xlsx":
		inputData, err = readDataFromXLSX(input)
	default:
		l.Fatalln("input file must be of type [.zip | .xls | .xlsx], got:", inputExt)
	}
	if err != nil {
		l.Fatal(err)
	}

	var genOptions []profile.GeneratorOption
	genOptions = append(
		genOptions,
		profile.WithGenerationTimestamp(*timestamp),
		profile.WithLogger(l),
	)
	if *sdkOverride != "" {
		genOptions = append(genOptions, profile.WithSDKVersionOverride(*sdkOverride))
	}
	if *switches {
		genOptions = append(genOptions, profile.WithSDKVersionOverride(*sdkOverride))
	}

	generator, err := profile.NewGenerator(input, inputData, genOptions...)
	if err != nil {
		l.Fatal(err)
	}

	fitProfile, err := generator.GenerateProfile()
	if err != nil {
		l.Fatal(err)
	}

	if err = ioutil.WriteFile(typesOut, fitProfile.TypesSource, 0644); err != nil {
		l.Fatalf("typegen: error writing types output file: %v", err)
	}

	if err = ioutil.WriteFile(messagesOut, fitProfile.MessagesSource, 0644); err != nil {
		l.Fatalf("typegen: error writing messages output file: %v", err)
	}

	if err = ioutil.WriteFile(profileOut, fitProfile.ProfileSource, 0644); err != nil {
		l.Fatalf("typegen: error writing profile output file: %v", err)
	}

	l.Println("running stringer")
	err = runStringerOnTypes(stringerPath, fitSrcDir, typesStringOut, fitProfile.StringerInput)
	if err != nil {
		l.Fatal(err)
	}
	l.Println("stringer: types done")

	err = runAllTests(fitPkgImportPath)
	if err != nil {
		l.Fatal(err)
	}
	l.Println("go test: pass")

	logMesgNumVsMessages(fitProfile.MesgNumsWithoutMessage, l)

	l.Println("done")
}

func runStringerOnTypes(stringerPath, fitSrcDir, typesStringOut, fitTypes string) error {

	stringerCmd := exec.Command(
		"go",
		"run",
		stringerPath,
		"-trimprefix",
		"-type", fitTypes,
		"-output",
		typesStringOut,
		fitSrcDir,
	)

	output, err := stringerCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("stringer: error running on types: %v\n%s", err, output)
	}

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
	// Packages
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

	return nil
}

func logMesgNumVsMessages(msgs []string, l *log.Logger) {
	if len(msgs) == 0 {
		return
	}
	l.Println("mesgnum-vs-msgs: implementation detail below, this may be automated in the future")
	l.Println("mesgnum-vs-msgs: #mesgnum values != #generated messages, diff:", len(msgs))
	l.Println("mesgnum-vs-msgs: remember to verify map in mappings.go for the following message(s):")
	for _, msg := range msgs {
		l.Printf("mesgnum-vs-msgs: ----> mesgnum %q has no corresponding message\n", msg)
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

func readDataFromZIP(path string) ([]byte, error) {
	r, err := zip.OpenReader(path)
	if err != nil {
		return nil, fmt.Errorf("error opening sdk zip file: %v", err)
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
		return nil, fmt.Errorf(
			"no file named %q or %q found in zip archive",
			workbookNameXLS, workbookNameXLSX,
		)
	}

	rc, err := wfile.Open()
	if err != nil {
		return nil, fmt.Errorf("error opening zip archive: %v", err)
	}
	defer rc.Close()

	data, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, fmt.Errorf("error reading %q from archive: %v", wfile.Name, err)
	}

	return data, nil
}

func readDataFromXLSX(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}
