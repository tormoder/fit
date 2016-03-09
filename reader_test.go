package fit_test

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/tormoder/fit"
)

func TestMain(m *testing.M) {
	err := parseActivitySmallData()
	if err != nil {
		fmt.Println("parseActivitySmallData failed:", err)
		os.Exit(2)
	}
	os.Exit(m.Run())
}

func TestDecode(t *testing.T) {
	for _, file := range testFiles {
		fpath := filepath.Join(tfolder, file.folder, file.name)
		t.Logf("decoding %q", fpath)

		data, err := ioutil.ReadFile(fpath)
		if err != nil {
			t.Fatalf("%q: reading file failed: %v", fpath, err)
		}

		fitFile, err := fit.Decode(bytes.NewReader(data))
		if !file.wantErr && err != nil {
			t.Errorf("%q: got error %v, want no error", fpath, err)
		}
		if file.wantErr && err == nil {
			t.Errorf("%q: got no error, want error", fpath)
		}

		if file.gobSHA1 == "" {
			continue
		}

		sum := sha1.New()
		enc := gob.NewEncoder(sum)
		err = enc.Encode(fitFile)
		if err != nil {
			t.Fatalf("%q: gob encode failed:", err)
		}
		b64sum := base64.StdEncoding.EncodeToString(sum.Sum(nil))
		if b64sum != file.gobSHA1 {
			t.Errorf("%q: SHA-1 for gob encoded fit file differs", fpath)
			// TODO(tormoder): Diff using goon?
		}
	}
}

func TestCheckIntegrity(t *testing.T) {
	// Implicitly tested for all files in TestDecode.
	// One example here.
	err := fit.CheckIntegrity(bytes.NewReader(activitySmallData), false)
	if err != nil {
		t.Errorf("%q: CheckIntegrity failed: %v", activitySmall, err)
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
	gotHeader, err := fit.DecodeHeader(bytes.NewReader(activitySmallData))
	if err != nil {
		t.Errorf("%q: DecodeHeader failed: %v", activitySmall, err)
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

	gotHeader, gotFileId, err := fit.DecodeHeaderAndFileID(bytes.NewReader(activitySmallData))
	if err != nil {
		t.Errorf("%q: DecodeHeaderAndFileId failed: %v", activitySmall, err)
	}
	if gotHeader != wantHeader {
		t.Errorf("%q:\ngot header:\n%#v\nwant header:\n%#v", activitySmall, gotHeader, wantHeader)
	}
	if gotFileId != wantFileId {
		t.Errorf("%q:\ngot FileIdMsg:\n%v\nwant FileIdMsg:\n%v", activitySmall, gotFileId, wantFileId)
	}
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
	data, err := ioutil.ReadFile(activityLarge)
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

func benchmarkDecode(b *testing.B, filename string, bench string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		b.Fatal(err)
	}
	b.ReportAllocs()
	b.SetBytes(int64(len(data)))
	switch bench {
	case "Full":
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, err := fit.Decode(bytes.NewReader(data))
			if err != nil {
				b.Fatal(err)
			}
		}
	case "Header":
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, err := fit.DecodeHeader(bytes.NewReader(data))
			if err != nil {
				b.Fatal(err)
			}
		}
	case "HeaderAndFileID":
		b.ResetTimer()
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

const (
	tfolder = "testdata" // root

	me       = "me"
	fitsdk   = "fitsdk"
	fitparse = "python-fitparse"
	sram     = "sram"
	dcrain   = "dcrainmaker"
	misc     = "misc"
	corrupt  = "corrupt"
)

var (
	activitySmall      = filepath.Join(tfolder, me, "activity-small-fenix2-run.fit")
	activityLarge      = filepath.Join(tfolder, me, "activity-large-fenxi2-multisport.fit")
	activityComponents = filepath.Join(tfolder, dcrain, "Edge810-Vector-2013-08-16-15-35-10.fit")
	activitySmallData  []byte
)

func parseActivitySmallData() error {
	var err error
	activitySmallData, err = ioutil.ReadFile(activitySmall)
	return err
}

var testFiles = [...]struct {
	folder  string
	name    string
	wantErr bool
	gobSHA1 string
}{
	{
		me,
		"activity-small-fenix2-run.fit",
		false,
		"",
	},
	{
		fitsdk,
		"Activity.fit",
		false,
		"",
	},
	{
		fitsdk,
		"MonitoringFile.fit",
		false,
		"",
	},
	{
		fitsdk,
		"Settings.fit",
		false,
		"",
	},

	{
		fitsdk,
		"WeightScaleMultiUser.fit",
		false,
		"",
	},
	{
		fitsdk,
		"WorkoutCustomTargetValues.fit",
		false,
		"",
	},
	{
		fitsdk,
		"WorkoutIndividualSteps.fit",
		false,
		"",
	},
	{
		fitsdk,
		"WorkoutRepeatGreaterThanStep.fit",
		false,
		"",
	},
	{
		fitsdk,
		"WorkoutRepeatSteps.fit",
		false,
		"",
	},
	{
		fitparse,
		"garmin-edge-500-activitiy.fit",
		false,
		"",
	},
	{
		fitparse,
		"sample-activity-indoor-trainer.fit",
		false,
		"",
	},
	{
		fitparse,
		"compressed-speed-distance.fit",
		false,
		"",
	},
	{
		fitparse,
		"antfs-dump.63.fit",
		false,
		"",
	},
	{
		sram,
		"Settings.fit",
		false,
		"",
	},
	{
		sram,
		"Settings2.fit",
		false,
		"",
	},
	{
		dcrain,
		"Edge810-Vector-2013-08-16-15-35-10.fit",
		false,
		"",
	},
	{
		misc,
		"2013-02-06-12-11-14.fit",
		false,
		"",
	},
	{
		misc,
		"2015-10-13-08-43-15.fit",
		false,
		"",
	},
	{
		corrupt,
		"activity-filecrc.fit",
		true,
		"",
	},
	{
		corrupt,
		"activity-unexpected-eof.fit",
		true,
		"",
	},
	/*
		{
		// TODO(tormoder): Investigate why this file fails.
		// "Decode failed for "testdata/fitsdk/WeightScaleSingleUser.fit": parsing
		// data message: missing data definition message for local message number
		// 1."
		fitsdk,
		"WeightScaleSingleUser.fit",
		false,
		"",
		},
	*/
}
