package main

import (
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

func main() {
	log.SetFlags(0)
	log.SetPrefix("fitgen:\t")

	fitSrcDir, err := goPackagePath(fitPkgImportPath)
	if err != nil {
		log.Fatalf("can't find fit package root src directory for %q", fitPkgImportPath)
	}
	log.Println("root src directory:", fitSrcDir)

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

	input := flag.Arg(0)
	inputExt := filepath.Ext(input)
	switch inputExt {
	case ".zip", ".xls", ".xlsx":
	default:
		log.Fatalln("input file must be of type [.zip | .xls | .xlsx], got:", inputExt)
	}

	var genOptions []profile.GeneratorOption
	genOptions = append(genOptions, profile.WithGenerationTimestamp(*timestamp))
	if *sdkOverride != "" {
		genOptions = append(genOptions, profile.WithSDKVersionOverride(*sdkOverride))
	}
	if *switches {
		genOptions = append(genOptions, profile.WithSDKVersionOverride(*sdkOverride))
	}

	generator, err := profile.NewGenerator(input, genOptions...)
	if err != nil {
		log.Fatal(err)
	}

	fitProfile, err := generator.GenerateProfile()
	if err != nil {
		log.Fatal(err)
	}

	if err = ioutil.WriteFile(typesOut, fitProfile.TypesSource, 0644); err != nil {
		log.Fatalf("typegen: error writing types output file: %v", err)
	}

	if err = ioutil.WriteFile(messagesOut, fitProfile.MessagesSource, 0644); err != nil {
		log.Fatalf("typegen: error writing messages output file: %v", err)
	}

	if err = ioutil.WriteFile(profileOut, fitProfile.ProfileSource, 0644); err != nil {
		log.Fatalf("typegen: error writing profile output file: %v", err)
	}

	err = runStringerOnTypes(stringerPath, fitSrcDir, typesStringOut, fitProfile.StringerInput)
	if err != nil {
		log.Fatal(err)
	}

	err = runAllTests(fitPkgImportPath)
	if err != nil {
		log.Fatal(err)
	}

	logMesgNumVsMessages(fitProfile.MesgNumsWithoutMessage)

	log.Println("done")
}

func runStringerOnTypes(stringerPath, fitSrcDir, typesStringOut, fitTypes string) error {
	log.Println("running stringer")

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

	log.Println("go test: pass")
	return nil
}

func logMesgNumVsMessages(msgs []string) {
	if len(msgs) == 0 {
		return
	}
	log.Println("mesgnum-vs-msgs: implementation detail below, this may be automated in the future")
	log.Println("mesgnum-vs-msgs: #mesgnum values != #generated messages, diff:", len(msgs))
	log.Println("mesgnum-vs-msgs: remember to verify map in mappings.go for the following message(s):")
	for _, msg := range msgs {
		log.Printf("mesgnum-vs-msgs: ----> mesgnum %q has no corresponding message\n", msg)
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
