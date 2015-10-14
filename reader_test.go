package fit_test

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/tormoder/fit"
)

func TestDecodeHeaderActivity(t *testing.T) {
	testDecode(t, activitySmall, true, false, false)
}

func TestDecodeHeaderAndFileID(t *testing.T) {
	testDecode(t, activitySmall, false, true, false)
}

func TestDecodeActivity(t *testing.T) {
	testDecode(t, activitySmall, false, false, false)
}

func TestCheckIntegrity(t *testing.T) {
	testDecode(t, activitySmall, false, false, true)
}

func TestFitSdkExamples(t *testing.T) {
	for _, test := range testsFitSDK {
		file := filepath.Join(tfolder, fitsdk, test)
		testDecode(t, file, false, false, false)
	}
}

func TestPytonFitparseTestdata(t *testing.T) {
	for _, test := range testsFitparse {
		file := filepath.Join(tfolder, fitparse, test)
		testDecode(t, file, false, false, false)
	}
}

func TestSramExamples(t *testing.T) {
	for _, test := range testsSram {
		file := filepath.Join(tfolder, sram, test)
		testDecode(t, file, false, false, false)
	}
}

func TestDCRainmakerFiles(t *testing.T) {
	for _, test := range testsDCRain {
		file := filepath.Join(tfolder, dcrain, test)
		testDecode(t, file, false, false, false)
	}
}

func TestMiscFitFilesFromWWW(t *testing.T) {
	for _, test := range testsMisc {
		file := filepath.Join(tfolder, misc, test)
		testDecode(t, file, false, false, false)
	}
}

func TestCorruptActivities(t *testing.T) {
	for _, test := range testsCorrupt {
		file := filepath.Join(tfolder, corrupt, test)
		err := testDecodeErr(t, file)
		if err == nil {
			t.Errorf("file: %q - want decoding error, got none", file)
		}
	}
}

func testDecode(t *testing.T, filename string, headerOnly, fileIDOnly, crcOnly bool) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	switch {
	case headerOnly:
		_, err = fit.DecodeHeader(bytes.NewReader(data))
		if err != nil {
			t.Errorf("DecodeHeader failed for %q: %v", filename, err)
		}
	case fileIDOnly:
		_, _, err = fit.DecodeHeaderAndFileID(bytes.NewReader(data))
		if err != nil {
			t.Errorf("DecodeHeaderAndFileID failed for %q: %v", filename, err)
		}
	case crcOnly:
		err = fit.CheckIntegrity(bytes.NewReader(data), false)
		if err != nil {
			t.Errorf("CheckIntegrity failed for %q: %v", filename, err)
		}
	default:
		t.Logf("Decoding %q", filename)
		_, err = fit.Decode(bytes.NewReader(data))
		if err != nil {
			t.Errorf("Decode failed for %q: %v", filename, err)
		}
	}
}

func testDecodeErr(t *testing.T, filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	_, err = fit.Decode(bytes.NewReader(data))
	return err
}

func BenchmarkDecodeActivitySmall(b *testing.B) {
	benchmarkDecode(b, activitySmall, "Full")
}

func BenchmarkDecodeActivityLarge(b *testing.B) {
	benchmarkDecode(b, activityLarge, "Full")
}

func BenchmarkDecodeActivityWithComponents(b *testing.B) {
	benchmarkDecode(b, activityComponents, "Full")
}

func BenchmarkDecodeHeader(b *testing.B) {
	benchmarkDecode(b, activitySmall, "Header")
}

func BenchmarkDecodeHeaderAndFileID(b *testing.B) {
	benchmarkDecode(b, activitySmall, "HeaderAndFileID")
}

func BenchmarkDecodeActivityLargeParallel(b *testing.B) {
	b.ReportAllocs()
	data, err := ioutil.ReadFile(activityLarge)
	if err != nil {
		b.Fatal(err)
	}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := fit.Decode(bytes.NewReader(data))
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func benchmarkDecode(b *testing.B, filename string, bench string) {
	b.ReportAllocs()
	b.StopTimer()
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		b.Fatal(err)
	}
	switch bench {
	case "Full":
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			_, err := fit.Decode(bytes.NewReader(data))
			if err != nil {
				b.Fatal(err)
			}
		}
	case "Header":
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			_, err := fit.DecodeHeader(bytes.NewReader(data))
			if err != nil {
				b.Fatal(err)
			}
		}
	case "HeaderAndFileID":
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			_, _, err := fit.DecodeHeaderAndFileID(bytes.NewReader(data))
			if err != nil {
				b.Fatal(err)
			}
		}
	default:
		panic("benchmarkDecode: unknown benchmark")
	}
}

// Testdata folders.
const (
	tfolder  = "testdata"
	me       = "me"
	fitsdk   = "fitsdk"
	fitparse = "python-fitparse"
	sram     = "sram"
	dcrain   = "dcrainmaker"
	misc     = "misc"
	corrupt  = "corrupt"
)

// Baseline activities.
var (
	activitySmall      = filepath.Join(tfolder, me, "activity-small-fenix2-run.fit")
	activityLarge      = filepath.Join(tfolder, me, "activity-large-fenxi2-multisport.fit")
	activityComponents = filepath.Join(tfolder, dcrain, testsDCRain[0])
)

// Example files from FIT SDK.
var testsFitSDK = []string{
	"Activity.fit",
	"MonitoringFile.fit",
	"Settings.fit",
	"WeightScaleMultiUser.fit",
	"WorkoutCustomTargetValues.fit",
	"WorkoutIndividualSteps.fit",
	"WorkoutRepeatGreaterThanStep.fit",
	"WorkoutRepeatSteps.fit",

	// TODO: Investigate why this file fails.
	/*
		"Decode failed for "testdata/fitsdk/WeightScaleSingleUser.fit": parsing
		data message: missing data definition message for local message number
		1."
	*/
	// "WeightScaleSingleUser.fit",
}

// Test data from python-fitparse and forks.
var testsFitparse = []string{
	"garmin-edge-500-activitiy.fit",
	"sample-activity-indoor-trainer.fit",
	"compressed-speed-distance.fit",
	"antfs-dump.63.fit",
}

// Examples from fit_json (Sram/Quarq)
var testsSram = []string{
	"Settings.fit",

	/*
		Earlier error for Settings2.fit:

		"Decode failed for "testdata/sram/Settings2.fit": parsing definition
		message: validating HrmProfile failed: base type for field 0 was
		uint8, but profile list field as of type byte"

		Has wrong base type for one profile field. Maybe encoded
		version of Settings.fit using fit_json, and this is an encoding bug?
		We now allow compatible types.
	*/
	"Settings2.fit",
}

// Power data published by DCRainmaker in power meter reviews.
var testsDCRain = []string{
	"Edge810-Vector-2013-08-16-15-35-10.fit",
}

// Miscellaneous fit files found online.
var testsMisc = []string{
	"2013-02-06-12-11-14.fit",
	"2015-10-13-08-43-15.fit",
}

// Intentionally corrupted fit files.
var testsCorrupt = []string{
	"activity-filecrc.fit",
	"activity-unexpected-eof.fit",
}
