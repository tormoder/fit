package profile

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

const (
	typesSheetIndex = 0
	msgsSheetIndex  = 1
)

func parseWorkbook(inputData []byte) (typeData, msgData [][]string, err error) {
	workbook, err := xlsx.OpenBinary(inputData)
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
