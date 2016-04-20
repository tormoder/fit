package profile

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/tealeg/xlsx"
)

func parseSDKVersionFromZipFile(zipFilePath string) string {
	// Brittle.
	// TODO: Maybe parse 'c/fit.h' with regexp instead.
	_, file := filepath.Split(zipFilePath)
	ver := strings.TrimSuffix(file, ".zip")
	return strings.TrimPrefix(ver, "FitSDKRelease_")
}

const (
	workbookNameXLS  = "Profile.xls"
	workbookNameXLSX = "Profile.xlsx"
	typesSheetIndex  = 0
	msgsSheetIndex   = 1
)

func parseWorkbook(inputPath string) (typeData, msgData [][]string, err error) {
	r, err := zip.OpenReader(inputPath)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening sdk zip file: %v", err)
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
		return nil, nil, fmt.Errorf(
			"no file named %q or %q found in zip archive",
			workbookNameXLS, workbookNameXLSX)
	}

	rc, err := wfile.Open()
	if err != nil {
		return nil, nil, err
	}
	defer rc.Close()
	b, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, nil, err
	}

	workbook, err := xlsx.OpenBinary(b)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening profile workbook: %v", err)
	}

	// file.ToSlice from the xlsx library adjusted to ignore formatting errors.
	var output = [][][]string{}
	for _, sheet := range workbook.Sheets {
		s := [][]string{}
		for _, row := range sheet.Rows {
			if row == nil {
				continue
			}
			r := []string{}
			for _, cell := range row.Cells {
				str, err := cell.String()
				if err != nil {
					// The profile message sheet has formatting errors.
					// Ignore those cells and use the raw values.
				}
				r = append(r, str)
			}
			s = append(s, r)
		}
		output = append(output, s)
	}

	typeData = output[typesSheetIndex]
	msgData = output[msgsSheetIndex]

	debugln("parse workbook: done")

	return typeData, msgData, nil
}
