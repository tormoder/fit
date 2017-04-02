package fit_test

import (
	"bytes"
	"compress/gzip"
	"flag"
	"fmt"
	"hash/fnv"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"

	"github.com/kortschak/utter"
	"github.com/tormoder/fit"
)

var update = flag.Bool("update", false, "update .golden output for decode test files")

func init() { flag.Parse() }

func fitFingerprint(fit *fit.Fit) uint32 {
	h := fnv.New32a()
	utter.Fdump(h, fit)
	return h.Sum32()
}

func fitUtterDump(fit *fit.Fit, path string, compressed bool) error {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	var w io.WriteCloser
	if compressed {
		w = gzip.NewWriter(f)
	} else {
		w = f
	}

	utter.Fdump(w, fit)

	if !compressed {
		return f.Close()
	}

	err = w.Close()
	if err != nil {
		_ = f.Close()
		return err
	}
	return f.Close()
}

var (
	loggerMu      sync.Mutex
	loggerOnce    sync.Once
	devNullLogger *log.Logger
)

func nullLogger() *log.Logger {
	loggerMu.Lock()
	defer loggerMu.Unlock()
	loggerOnce.Do(func() {
		devNullLogger = log.New(ioutil.Discard, "", 0)
	})
	return devNullLogger
}

var (
	activitySmallMu   sync.Mutex
	activitySmallOnce sync.Once
	activitySmallData []byte
)

func activitySmall() []byte {
	activitySmallMu.Lock()
	defer activitySmallMu.Unlock()
	activitySmallOnce.Do(func() {
		asd, err := ioutil.ReadFile(activitySmallPath)
		if err != nil {
			errDesc := fmt.Sprintf("parseActivitySmallData failed: %v", err)
			panic(errDesc)
		}
		activitySmallData = asd
	})
	return activitySmallData
}

var (
	activitySmallPath      = filepath.Join(tfolder, me, "activity-small-fenix2-run.fit")
	activityLargePath      = filepath.Join(tfolder, me, "activity-large-fenxi2-multisport.fit")
	activityComponentsPath = filepath.Join(tfolder, dcrain, "Edge810-Vector-2013-08-16-15-35-10.fit")
)

const (
	goldenSuffix  = ".golden"
	currentSuffix = ".current"
	gzSuffix      = ".gz"
)

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

func TestDecode(t *testing.T) {
	testFiles := [...]struct {
		folder      string
		name        string
		wantErr     bool
		fingerprint uint32
		compress    bool
		dopts       []fit.DecodeOption
	}{
		{
			me,
			"activity-small-fenix2-run.fit",
			false,
			389865072,
			true,
			[]fit.DecodeOption{
				fit.WithUnknownFields(),
				fit.WithUnknownMessages(),
				fit.WithLogger(nullLogger()), // For test coverage.
			},
		},
		{
			fitsdk,
			"Activity.fit",
			false,
			2722395800,
			true,
			nil,
		},
		{
			fitsdk,
			"MonitoringFile.fit",
			false,
			1628070551,
			true,
			nil,
		},
		{
			fitsdk,
			"Settings.fit",
			false,
			300677984,
			true,
			nil,
		},

		{
			fitsdk,
			"WeightScaleMultiUser.fit",
			false,
			1631563326,
			true,
			nil,
		},
		{
			fitsdk,
			"WorkoutCustomTargetValues.fit",
			false,
			407208810,
			true,
			nil,
		},
		{
			fitsdk,
			"WorkoutIndividualSteps.fit",
			false,
			1159208415,
			true,
			nil,
		},
		{
			fitsdk,
			"WorkoutRepeatGreaterThanStep.fit",
			false,
			3635471829,
			true,
			nil,
		},
		{
			fitsdk,
			"WorkoutRepeatSteps.fit",
			false,
			2404845921,
			true,
			nil,
		},
		{
			fitparse,
			"garmin-edge-500-activitiy.fit",
			false,
			0,
			false,
			nil,
		},
		{
			fitparse,
			"sample-activity-indoor-trainer.fit",
			false,
			0,
			false,
			nil,
		},
		{
			fitparse,
			"compressed-speed-distance.fit",
			false,
			0,
			false,
			nil,
		},
		{
			fitparse,
			"antfs-dump.63.fit",
			false,
			0,
			false,
			nil,
		},
		{
			sram,
			"Settings.fit",
			false,
			0,
			false,
			nil,
		},
		{
			sram,
			"Settings2.fit",
			false,
			0,
			false,
			nil,
		},
		{
			dcrain,
			"Edge810-Vector-2013-08-16-15-35-10.fit",
			false,
			0,
			false,
			nil,
		},
		{
			misc,
			"2013-02-06-12-11-14.fit",
			false,
			0,
			false,
			nil,
		},
		{
			misc,
			"2015-10-13-08-43-15.fit",
			false,
			0,
			false,
			nil,
		},
		{
			corrupt,
			"activity-filecrc.fit",
			true,
			0,
			false,
			nil,
		},
		{
			corrupt,
			"activity-unexpected-eof.fit",
			true,
			0,
			false,
			nil,
		},
	}
	for _, file := range testFiles {
		file := file
		t.Run(fmt.Sprintf("%s/%s", file.folder, file.name), func(t *testing.T) {
			t.Parallel()
			fpath := filepath.Join(tfolder, file.folder, file.name)
			data, err := ioutil.ReadFile(fpath)
			if err != nil {
				t.Fatalf("reading file failed: %v", err)
			}
			fitFile, err := fit.Decode(bytes.NewReader(data), file.dopts...)
			if !file.wantErr && err != nil {
				t.Errorf("got error %v, want no error", err)
			}
			if file.wantErr && err == nil {
				t.Error("got no error, want error")
			}
			if file.fingerprint == 0 || file.wantErr {
				return
			}
			fp := fitFingerprint(fitFile)
			if fp == file.fingerprint {
				return
			}
			t.Errorf("fit file fingerprint differs: got: %d, want: %d", fp, file.fingerprint)
			if !*update {
				fpath = fpath + currentSuffix
			} else {
				fpath = fpath + goldenSuffix
			}
			if file.compress {
				fpath = fpath + gzSuffix
			}
			err = fitUtterDump(fitFile, fpath, file.compress)
			if err != nil {
				t.Fatalf("error writing output: %v", err)
			}
			if !*update {
				t.Logf("current output written to: %s", fpath)
				t.Logf("use a diff tool to compare (e.g. zdiff if compressed)")
			} else {
				t.Logf("%q has been updated", fpath)
				t.Logf("new fingerprint is: %d, update test case in reader_test.go", fp)
			}
		})
	}
}

func TestCheckIntegrity(t *testing.T) {
	// Implicitly tested for all files in TestDecode.
	// One example here.
	err := fit.CheckIntegrity(bytes.NewReader(activitySmall()), false)
	if err != nil {
		t.Errorf("%q: failed: %v", activitySmallPath, err)
	}
}

func TestDecodeHeader(t *testing.T) {
	wantHeader := fit.Header{
		Size:            0xe,
		ProtocolVersion: 0x10,
		ProfileVersion:  0x457,
		DataSize:        0x1dbdf,
		DataType:        [4]uint8{0x2e, 0x46, 0x49, 0x54},
		CRC:             0x1ec4,
	}
	gotHeader, err := fit.DecodeHeader(bytes.NewReader(activitySmall()))
	if err != nil {
		t.Errorf("%q: failed: %v", activitySmallPath, err)
	}
	if gotHeader != wantHeader {
		t.Errorf("got header:\n%#v\nwant header: %#v", gotHeader, wantHeader)
	}
}

func TestDecodeHeaderAndFileID(t *testing.T) {
	wantHeader := fit.Header{
		Size:            0xe,
		ProtocolVersion: 0x10,
		ProfileVersion:  0x457,
		DataSize:        0x1dbdf,
		DataType:        [4]uint8{0x2e, 0x46, 0x49, 0x54},
		CRC:             0x1ec4,
	}
	tc := time.Unix(1439652761, 0)
	tc = tc.UTC()
	wantFileId := fit.FileIdMsg{
		Type:         0x4,
		Manufacturer: 0x1,
		Product:      0x7af,
		SerialNumber: 0xe762d9cf,
		Number:       0xffff,
		TimeCreated:  tc,
		ProductName:  "",
	}

	gotHeader, gotFileId, err := fit.DecodeHeaderAndFileID(bytes.NewReader(activitySmall()))
	if err != nil {
		t.Errorf("%q: failed: %v", activitySmallPath, err)
	}
	if gotHeader != wantHeader {
		t.Errorf("%q:\ngot header:\n%#v\nwant header:\n%#v", activitySmallPath, gotHeader, wantHeader)
	}
	if gotFileId != wantFileId {
		t.Errorf("%q:\ngot FileIdMsg:\n%v\nwant FileIdMsg:\n%v", activitySmallPath, gotFileId, wantFileId)
	}
}

func BenchmarkDecodeActivity(b *testing.B) {
	files := []struct {
		desc, path string
	}{
		{"Small", activitySmallPath},
		{"Large", activityLargePath},
		{"WithComponents", activityComponentsPath},
	}
	for _, file := range files {
		b.Run(file.desc, func(b *testing.B) {
			data, err := ioutil.ReadFile(file.path)
			if err != nil {
				b.Fatalf("%q: error reading file: %v", file.path, err)
			}
			b.ReportAllocs()
			b.SetBytes(int64(len(data)))
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, err := fit.Decode(bytes.NewReader(data))
				if err != nil {
					b.Fatalf("%q: error decoding file: %v", file.path, err)
				}
			}
		})
	}
}

func BenchmarkDecodeActivityLargeParallel(b *testing.B) {
	data, err := ioutil.ReadFile(activityLargePath)
	if err != nil {
		b.Fatal(err)
	}
	b.ReportAllocs()
	b.SetBytes(int64(len(data)))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := fit.Decode(bytes.NewReader(data))
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkDecodeHeader(b *testing.B) {
	data := activitySmall()
	b.ReportAllocs()
	b.SetBytes(int64(len(data)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := fit.DecodeHeader(bytes.NewReader(data))
		if err != nil {
			b.Fatalf("%q: error decoding header: %v", activitySmallPath, err)
		}
	}

}

func BenchmarkDecodeHeaderAndFileID(b *testing.B) {
	data := activitySmall()
	b.ReportAllocs()
	b.SetBytes(int64(len(data)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, err := fit.DecodeHeaderAndFileID(bytes.NewReader(data))
		if err != nil {
			b.Fatalf("%q: error decoding header/fileid: %v", activitySmallPath, err)
		}
	}
}
