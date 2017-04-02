package profile_test

import (
	"bytes"
	"flag"
	"io"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/cespare/xxhash"
	"github.com/tormoder/fit/cmd/fitgen/internal/profile"
)

const (
	testdata      = "testdata"
	fileExt       = ".xlsx"
	goldenSuffix  = ".golden"
	currentSuffix = ".current"
)

var update = flag.Bool("update", false, "update .golden output files")

func init() { flag.Parse() }

var currentSDK = sdks[0]

var defGenOpts = []profile.GeneratorOption{
	profile.WithGenerationTimestamp(false),
}

func relPath(sdkVersion string) string {
	return filepath.Join(testdata, sdkVersion+fileExt)
}

func writeProfile(p *profile.Profile, w io.Writer) error {
	var err error
	write := func(buf []byte) {
		if err != nil {
			return
		}
		_, err = w.Write(buf)
	}
	write([]byte("// TYPES\n"))
	write(p.TypesSource)
	write([]byte("// MESSAGES\n"))
	write(p.MessagesSource)
	write([]byte("// PROFILE\n"))
	write(p.ProfileSource)
	write([]byte("// STRINGER TYPE INPUT\n"))
	write([]byte(p.StringerInput))
	write([]byte("\n// MESSAGE NUMS WITHOUT MESSAGE\n"))
	for _, mn := range p.MesgNumsWithoutMessage {
		write([]byte(mn))
		write([]byte{'\n'})
	}
	return err
}

func writeProfileToFile(p *profile.Profile, path string) error {
	buf := new(bytes.Buffer)
	err := writeProfile(p, buf)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, buf.Bytes(), 0644)
}

func profileFingerprint(p *profile.Profile) uint64 {
	h := xxhash.New()
	_ = writeProfile(p, h)
	return h.Sum64()
}

type sdk struct {
	version           string
	goldenFingerprint uint64
}

var sdks = []sdk{
	{"16.20", 4528042509788505848},
	{"20.14", 4959578396414666089},
}

func TestGenerator(t *testing.T) {
	for _, sdk := range sdks {
		t.Run(sdk.version, func(t *testing.T) {
			if sdk == currentSDK && testing.Short() {
				t.Skip("skipping test in short mode")
			}
			path := relPath(sdk.version)
			data, err := ioutil.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			g, err := profile.NewGenerator(path, data, defGenOpts...)
			if err != nil {
				t.Fatal(err)
			}
			p, err := g.GenerateProfile()
			if err != nil {
				t.Fatal(err)
			}
			gotFP := profileFingerprint(p)
			if gotFP == sdk.goldenFingerprint {
				return
			}
			t.Errorf("profile fingerprint differs: got: %d, want %d", gotFP, sdk.goldenFingerprint)
			if !*update {
				path = path + currentSuffix
			} else {
				path = path + goldenSuffix
			}
			err = writeProfileToFile(p, path)
			if err != nil {
				t.Fatalf("error writing output: %v", err)
			}
			if !*update {
				t.Logf("current output written to: %s", path)
			} else {
				t.Logf("%q has been updated", path)
				t.Logf("new fingerprint is: %d", gotFP)
			}
		})
	}
}

var profileSink *profile.Profile

func BenchmarkGenerator(b *testing.B) {
	for _, sdk := range sdks {
		b.Run(sdk.version, func(b *testing.B) {
			path := relPath(sdk.version)
			data, err := ioutil.ReadFile(path)
			if err != nil {
				b.Fatalf("error reading profile workbook: %v", err)
			}
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				g, err := profile.NewGenerator(path, data, defGenOpts...)
				if err != nil {
					b.Fatal(err)
				}
				profileSink, err = g.GenerateProfile()
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}
