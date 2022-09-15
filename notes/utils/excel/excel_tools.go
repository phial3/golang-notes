package excel

import (
	"bytes"
	"fmt"
	"io"
	"math"
)

import (
	"github.com/spf13/cast"
	"github.com/xuri/excelize/v2"
)

// ReadExcel filename: 指定文件路径名称，例如："/tmp/aa.xlsx"
// tableName: 可变参数,指定Excel中的某个表格名,不传默认第一个表格
func ReadExcel(filename string, tableName ...string) ([][]string, error) {
	f, err := excelize.OpenFile(filename) //
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	firstSheet := ""
	if len(tableName) > 0 {
		firstSheet = tableName[0]
	} else {
		firstSheet = f.GetSheetName(0)
	}
	rows, err := f.GetRows(firstSheet)
	return rows, err
}

// File is a struct for file
type File interface {
	io.Reader
	io.ReaderAt
	io.Seeker
	io.Closer
}

// Read form data bytes
func Read(dataBytes []byte, tableName ...string) ([][]string, error) {
	f, err := excelize.OpenReader(bytes.NewReader(dataBytes))
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	firstSheet := ""
	if len(tableName) > 0 {
		firstSheet = tableName[0]
	} else {
		firstSheet = f.GetSheetName(0)
	}
	rows, err := f.GetRows(firstSheet)
	return rows, err
}

// ReadExcelFile file stream
func ReadExcelFile(file File, tableName ...string) ([][]string, error) {
	f, err := excelize.OpenReader(file)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	//默认读取第一个
	firstSheet := ""
	if len(tableName) > 0 {
		firstSheet = tableName[0]
	} else {
		firstSheet = f.GetSheetName(0)
	}
	rows, err := f.GetRows(firstSheet)
	return rows, err
}

func WriteExcel(filename string, value [][]string, sheetName ...string) error {
	f := excelize.NewFile()
	//默认Excel每个页面保存500条数据，超出500条就新建一个页面保存，page为每页最多保存条数
	page := 1000
	//默认sheet名称为Sheet
	firstSheet := "Sheet"
	if len(sheetName) > 0 {
		firstSheet = sheetName[0]
	}

	// Create a new sheet.
	sheetRow := int(math.Floor(float64(len(value)/page)) + 1)
	for j := 0; j < sheetRow; j++ {
		index := f.NewSheet(firstSheet + cast.ToString(j+1))
		f.SetActiveSheet(index - 1)

		if j == 0 {
			for i := 0 + j*page; i < (j+1)*page; i++ { //行
				if len(value) < i+1 {
					break
				}
				for k, v := range value[i] { //列
					path, err := excelize.ColumnNumberToName(k + 1)
					if err != nil {
						return err
					}
					err = f.SetCellValue(firstSheet+cast.ToString(j+1), path+cast.ToString(i+1-j*page), v)
					if err != nil {
						return err
					}
				}
			}
			continue
		}

		//列
		for k, v := range value[0] {
			path, err := excelize.ColumnNumberToName(k + 1)
			if err != nil {
				return err
			}
			err = f.SetCellValue(firstSheet+cast.ToString(j+1), path+cast.ToString(1), v)
			if err != nil {
				return err
			}
		}

		//行
		for i := 0 + j*page; i < (j+1)*page; i++ {
			if len(value) < i+1 {
				break
			}
			for k, v := range value[i] { //列
				path, err := excelize.ColumnNumberToName(k + 1)
				if err != nil {
					return err
				}
				err = f.SetCellValue(firstSheet+cast.ToString(j+1), path+cast.ToString(i+2-j*page), v)
				if err != nil {
					return err
				}
			}
		}

	}
	if err := f.SaveAs(filename); err != nil {
		return err
	}
	return nil
}
