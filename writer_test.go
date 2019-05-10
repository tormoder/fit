package fit_test

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/tormoder/fit"
)

func TestDecodeEncodeDecode(t *testing.T) {
	t.Run("Group", func(t *testing.T) {
		for i, file := range decodeTestFiles {
			_, file := i, file // Capture range variables.

			if file.wantErr || file.fingerprint == 0 || file.skipEncode {
				continue
			}

			t.Run(fmt.Sprintf("%s/%s", file.folder, file.name), func(t *testing.T) {
				t.Parallel()
				fpath := filepath.Join(tdfolder, file.folder, file.name)

				inData, err := ioutil.ReadFile(fpath)
				if err != nil {
					t.Fatalf("reading file failed: %v", err)
				}

				inFile, err := fit.Decode(bytes.NewReader(inData))
				if err != nil {
					t.Fatalf("decode: got error, want none; error is: %v", err)
				}

				// Sanity check that decoding is OK
				fp := fitFingerprint(inFile)

				if fp != file.fingerprint {
					t.Fatalf("decode: fit file fingerprint differs: got: %d, want: %d", fp, file.fingerprint)
				}

				outBuf := &bytes.Buffer{}

				err = fit.Encode(outBuf, inFile, binary.LittleEndian)
				if err != nil {
					t.Fatalf("encode: got error, want none; error is: %v", err)
				}

				reFile, err := fit.Decode(bytes.NewReader(outBuf.Bytes()))
				if err != nil {
					t.Fatalf("re-decode: got error, want none; error is: %v", err)
				}

				// Wipe the CRCs. The serialized data is likely to be different,
				// so we need to ignore the CRCs.
				inFile.CRC = 0
				inFile.Header.CRC = 0
				reFile.CRC = 0
				reFile.Header.CRC = 0

				// Re-fingerprint without the CRCs
				fp = fitFingerprint(inFile)

				refp := fitFingerprint(reFile)
				if refp != fp {
					t.Errorf("re-decode: fit file fingerprint differs: got: %d, want: %d", refp, fp)

					fpath = fpath + currentSuffix
					err = fitUtterDump(reFile, fpath, false)
					if err != nil {
						t.Fatalf("error writing output: %v", err)
					}
					t.Logf("current output written to: %s", fpath)

					fpath = fpath + ".in"
					err = fitUtterDump(inFile, fpath, false)
					if err != nil {
						t.Fatalf("error writing output: %v", err)
					}
					t.Logf("current input written to: %s", fpath)
				}
			})
		}
	})
}
