package profile_test

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
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

func scanLinesPreserveEOL(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.IndexByte(data, '\n'); i >= 0 {
		// We have a full newline-terminated line.
		return i + 1, data[0 : i+1], nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), data, nil
	}
	// Request more data.
	return 0, nil, nil
}

func readGoldenProfile(path string) (*profile.Profile, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() { _ = f.Close() }() // Sigh. To keep errcheck happy

	p := &profile.Profile{}
	headings := []string{
		"// TYPES",
		"// MESSAGES",
		"// PROFILE",
		"// STRINGER TYPE INPUT",
		"// MESSAGE NUMS WITHOUT MESSAGE",
	}
	i := 0

	scanner := bufio.NewScanner(f)
	scanner.Split(scanLinesPreserveEOL)

	scanner.Scan()
	err = scanner.Err()
	if err != nil {
		return nil, err
	}

	if !strings.HasPrefix(scanner.Text(), headings[i]) {
		return nil, fmt.Errorf("first line should be '%s'", headings[i])
	}

	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), headings[i+1]) {
			i++
			break
		}
		p.TypesSource = append(p.TypesSource, scanner.Bytes()...)
	}

	// Format
	p.TypesSource, err = format.Source(p.TypesSource)
	if err != nil {
		return nil, err
	}

	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), headings[i+1]) {
			i++
			break
		}
		p.MessagesSource = append(p.MessagesSource, scanner.Bytes()...)
	}

	// Format
	p.MessagesSource, err = format.Source(p.MessagesSource)
	if err != nil {
		return nil, err
	}

	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), headings[i+1]) {
			i++
			break
		}
		p.ProfileSource = append(p.ProfileSource, scanner.Bytes()...)
	}

	// Format
	p.ProfileSource, err = format.Source(p.ProfileSource)
	if err != nil {
		return nil, err
	}

	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), headings[i+1]) {
			break
		}
		p.StringerInput += strings.TrimSpace(scanner.Text())
	}

	for scanner.Scan() {
		p.MesgNumsWithoutMessage = append(p.MesgNumsWithoutMessage, strings.TrimSpace(scanner.Text()))
	}

	return p, scanner.Err()
}

func profileFingerprint(p *profile.Profile) uint64 {
	h := xxhash.New()
	_ = writeProfile(p, h)
	return h.Sum64()
}

type sdk struct {
	majVer, minVer int
}

var sdks = []sdk{
	{16, 20},
	{20, 14},
	{20, 27},
	{20, 43},
}

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}

func TestGenerator(t *testing.T) {
	for _, sdk := range sdks {
		sdk := sdk // Capture range variable.
		sdkFullVer := fmt.Sprintf("%d.%d", sdk.majVer, sdk.minVer)
		t.Run(sdkFullVer, func(t *testing.T) {
			t.Parallel()
			if sdk == currentSDK && testing.Short() {
				t.Skip("skipping test in short mode")
			}
			path := relPath(sdkFullVer)
			data, err := ioutil.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			g, err := profile.NewGenerator(sdk.majVer, sdk.minVer, data, defGenOpts...)
			if err != nil {
				t.Fatal(err)
			}
			p, err := g.GenerateProfile()
			if err != nil {
				t.Fatal(err)
			}
			gotFP := profileFingerprint(p)

			// Read in the golden profile, format it and fingerprint it
			// This makes the test robust against gofmt changes
			goldenProfile, err := readGoldenProfile(path + ".golden")
			if err != nil {
				t.Fatal(err)
			}
			goldenFingerprint := profileFingerprint(goldenProfile)

			if gotFP == goldenFingerprint {
				return
			}
			t.Errorf("profile fingerprint differs: got: %d, want %d", gotFP, goldenFingerprint)
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
		sdkFullVer := fmt.Sprintf("%d.%d", sdk.majVer, sdk.minVer)
		b.Run(sdkFullVer, func(b *testing.B) {
			path := relPath(sdkFullVer)
			data, err := ioutil.ReadFile(path)
			if err != nil {
				b.Fatalf("error reading profile workbook: %v", err)
			}
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				g, err := profile.NewGenerator(sdk.majVer, sdk.minVer, data, defGenOpts...)
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
