package helper

import (
	"boilerplate/exception"
	"fmt"
	"github.com/xuri/excelize/v2"
	"mime/multipart"
)

func removeEmptyStrings(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func ParseExcelToArray(file *multipart.FileHeader) ([]string, error) {
	theFile, err := file.Open()
	if err != nil {
		fmt.Println(err)
		return nil, exception.INVALID_EXCEL_FILE
	}

	excel, err := excelize.OpenReader(theFile)
	if err != nil {
		fmt.Println(err)

		return nil, exception.INVALID_EXCEL_FILE
	}

	sheetMap := excel.GetSheetMap()
	if len(sheetMap) == 0 {
		fmt.Println(sheetMap)
		return nil, exception.INVALID_EXCEL_FILE
	}

	rows, err := excel.GetCols(sheetMap[1])
	if err != nil {
		fmt.Println(sheetMap)
		return nil, exception.INVALID_EXCEL_FILE
	}
	if len(rows) == 0 {
		return nil, exception.INVALID_EXCEL_FILE
	}
	row := removeEmptyStrings(rows[0])
	if len(row) == 1 {
		return nil, exception.INVALID_EXCEL_FILE
	}
	return row[1:], nil
}
