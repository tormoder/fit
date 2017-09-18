package fit_test

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/tormoder/fit"
)

var (
	update  = flag.Bool("update", false, "update .golden output and table for decode test files if their fingerprint differs")
	fupdate = flag.Bool("fupdate", false, "force regeneration of decode test files table")
	fdecode = flag.Bool("fdecode", false, "force decode golden part of decode test irregardless of Go version")
)

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
	activitySmallPath      = filepath.Join(tdfolder, "me", "activity-small-fenix2-run.fit")
	activityLargePath      = filepath.Join(tdfolder, "me", "activity-large-fenxi2-multisport.fit")
	activityComponentsPath = filepath.Join(tdfolder, "dcrainmaker", "Edge810-Vector-2013-08-16-15-35-10.fit")
	monitoringPath         = filepath.Join(tdfolder, "fitsdk", "MonitoringFile.fit")
)

const (
	goldenSuffix  = ".golden"
	currentSuffix = ".current"
	gzSuffix      = ".gz"
	tdfolder      = "testdata"
)

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}

func TestDecode(t *testing.T) {
	const goMajorVersionForDecodeGolden = "go1.9"
	testDecodeGolden := true
	goVersion := runtime.Version()
	goVersionOK := strings.HasPrefix(goVersion, goMajorVersionForDecodeGolden)
	switch {
	case !goVersionOK && !*fdecode:
		testDecodeGolden = false
		t.Logf(
			"skipping golden decode part of test due to Go version (enabled for %s.x, have %q)",
			goMajorVersionForDecodeGolden,
			goVersion,
		)
	case !goVersionOK && *fdecode:
		t.Logf(
			"override: performing golden decode part of test for Go version %q (default only for %s.x)",
			goVersion,
			goMajorVersionForDecodeGolden,
		)
	default:
	}

	regenTestTable := struct {
		sync.Mutex // Protects val and decodeTestFiles slice in reader_util_test.go.
		val        bool
	}{}

	t.Run("Group", func(t *testing.T) {
		for i, file := range decodeTestFiles {
			i, file := i, file // Capture range variables.
			t.Run(fmt.Sprintf("%s/%s", file.folder, file.name), func(t *testing.T) {
				t.Parallel()
				fpath := filepath.Join(tdfolder, file.folder, file.name)
				data, err := ioutil.ReadFile(fpath)
				if err != nil {
					t.Fatalf("reading file failed: %v", err)
				}
				fitFile, err := fit.Decode(bytes.NewReader(data), file.dopts.opts()...)
				if !file.wantErr && err != nil {
					t.Fatalf("got error, want none; error is: %v", err)
				}
				if file.wantErr && err == nil {
					t.Fatalf("got no error, want error")
				}
				if !testDecodeGolden {
					return
				}
				if file.fingerprint == 0 {
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
					regenTestTable.Lock()
					regenTestTable.val = true
					decodeTestFiles[i].fingerprint = fp
					regenTestTable.Unlock()
					t.Logf("%q has been updated", fpath)
					t.Logf("new fingerprint is: %d, update test case in reader_test.go", fp)
				}
			})
		}
	})

	if regenTestTable.val || *fupdate {
		t.Logf("regenerating table for decode test files...")
		err := regenerateDecodeTestTable()
		if err != nil {
			t.Fatalf("error regenerating table for decode test files: %v", err)
		}
	}
}

func TestDecodeChained(t *testing.T) {
	chainedTestFiles := []struct {
		fpath   string
		dfiles  int
		wantErr bool
		desc    string
	}{
		{
			filepath.Join(tdfolder, "fitsdk", "Activity.fit"),
			1,
			false,
			"single valid fit file",
		},
		{
			filepath.Join(tdfolder, "chained", "activity-settings.fit"),
			2,
			false,
			"two valid chained fit files",
		},
		{
			filepath.Join(tdfolder, "chained", "activity-activity-filecrc.fit"),
			2,
			true,
			"one valid fit file + one fit file with wrong crc",
		},
		{
			filepath.Join(tdfolder, "chained", "activity-settings-corruptheader.fit"),
			1,
			true,
			"one valid fit file + one fit file with corrupt header",
		},
		{
			filepath.Join(tdfolder, "chained", "activity-settings-nodata.fit"),
			2,
			true,
			"one valid fit file + one fit file with ok header but no data",
		},
	}

	for _, ctf := range chainedTestFiles {
		ctf := ctf
		t.Run(ctf.fpath, func(t *testing.T) {
			t.Parallel()
			data, err := ioutil.ReadFile(ctf.fpath)
			if err != nil {
				t.Fatalf("reading file data failed: %v", err)
			}
			fitFiles, err := fit.DecodeChained(bytes.NewReader(data))
			if !ctf.wantErr && err != nil {
				t.Fatalf("got error, want none; error is: %v", err)
			}
			if ctf.wantErr && err == nil {
				t.Fatalf("got no error, want error")
			}
			if len(fitFiles) != ctf.dfiles {
				t.Fatalf("got %d decoded fit file(s), want %d", len(fitFiles), ctf.dfiles)
			}
		})
	}
}

func TestCheckIntegrity(t *testing.T) {
	t.Run("ActivitySmall", func(t *testing.T) {
		err := fit.CheckIntegrity(bytes.NewReader(activitySmall()), false)
		if err != nil {
			t.Errorf("%q: failed: %v", activitySmallPath, err)
		}
	})
	t.Run("ActivitySDK", func(t *testing.T) {
		fpath := filepath.Join(tdfolder, "fitsdk", "Activity.fit")
		data, err := ioutil.ReadFile(fpath)
		if err != nil {
			t.Fatalf("reading %q failed: %v", fpath, err)
		}
		err = fit.CheckIntegrity(bytes.NewReader(data), false)
		if err != nil {
			t.Errorf("%q: failed: %v", fpath, err)
		}
	})
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

func BenchmarkDecode(b *testing.B) {
	files := []struct {
		desc, path string
	}{
		{"ActivitySmall", activitySmallPath},
		{"ActivityLarge", activityLargePath},
		{"ActivityWithComponents", activityComponentsPath},
		{"MonitoringFile", monitoringPath},
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
