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
		ncols := len(sheet.Rows[0].Cells)
		for _, row := range sheet.Rows {
			if row == nil {
				continue
			}
			r := make([]string, ncols)
			for i, cell := range row.Cells {
				// The profile message sheet has formatting errors.
				// Ignore those cells and use the raw values.
				r[i] = cell.String()
			}
			s = append(s, r)
		}
		output = append(output, s)
	}

	typeData = output[typesSheetIndex]
	msgData = output[msgsSheetIndex]

	return typeData, msgData, nil
}
