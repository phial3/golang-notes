package excel

import (
	"fmt"
	"github.com/phial3/go-notes/utils/commonregex"
	"os"
	"testing"
)

import (
	"github.com/mitchellh/mapstructure"
)

// ExcelOfficeInfo 定义Excel中的一行数据
// 结构体TAG中 json,index 必须存在
// json: 字段名称
// index: 索引名称
// name: 中文名称
type ExcelOfficeInfo struct {
	OfficeCode              string `json:"officeCode" name:"OFFICE号" index:"0"`
	IataCode                string `json:"iataCode" name:"IATA号" index:"1"`
	OfficeName              string `json:"officeName" name:"代理人中文名称" index:"2"`
	OfficeNameEN            string `json:"officeNameEN" name:"代理人英文名称" index:"3"`
	City                    string `json:"city" name:"所在城市" index:"4"`
	Country                 string `json:"country" name:"所在国家" index:"5"`
	PayName                 string `json:"payName" name:"需修改为的结算名称" index:"6"`
	Description             string `json:"description" name:"备注" index:"7"`
	UnifiedSocialCreditCode string `json:"unifiedSocialCreditCode" name:"统一社会信用代码" index:"8"`
}

type BranchOfficeRelationship struct {
	OfficeCode string `json:"officeCode" name:"OFFICE号" index:"0"`
	OfficeName string `json:"officeName" name:"客户全称" index:"1"`
	BranchName string `json:"branchName" name:"分支" index:"2"`
	BranchCode string `json:"branchCode" name:"分支机构代码" index:"3"`
}

const (
	OfficeInfoExcelFileAbsolutePath string = "/tmp/5.billing_SLYX_AGENT_20220708.xlsx"
)

func TestPrintOfficeInfo(t *testing.T) {
	rows, err := ReadExcel(OfficeInfoExcelFileAbsolutePath)
	if err != nil {
		t.Fatal("read excel error. err=", err.Error())
	}

	for _, row := range rows {
		fmt.Println("row=", row)
	}
}

func TestParseStruct(t *testing.T) {
	rows, err := ReadExcel(OfficeInfoExcelFileAbsolutePath)
	if err != nil {
		t.Fatal("read excel error. err=", err.Error())
	}

	var arr []ExcelOfficeInfo
	err = NewExcelStructDefault().SetPointerStruct(&ExcelOfficeInfo{}).RowsAllProcess(rows,
		func(maps map[string]interface{}) error {
			var ptr ExcelOfficeInfo
			// map 转 结构体
			if err2 := mapstructure.Decode(maps, &ptr); err2 != nil {
				return err2
			}

			arr = append(arr, ptr)

			return nil
		})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	repeatMap := make(map[string]string, 0)
	// 不满足条件的过滤掉
	for idx, v := range arr {
		if v.OfficeCode != "" && v.OfficeName != "" && commonregex.HasZhFullChar(v.OfficeName) {
			_, ok := repeatMap[v.OfficeCode]
			if ok {
				fmt.Printf("重复数据 %d: %#v\n", idx, v)
			} else {
				repeatMap[v.OfficeCode] = v.OfficeName
			}
		}
	}
	fmt.Printf("总共%d条\n", len(repeatMap))

}

func TestWriteExcel(t *testing.T) {
	type args struct {
		filename  string
		value     [][]string
		tableName []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test_write_excel",
			args: args{
				filename: "/tmp/test_write_excel.xlsx",
				value: [][]string{
					[]string{"1", "2", "3"},
					[]string{"4", "5", "6"},
					[]string{"7", "8", "9"},
				},
				tableName: []string{"WR", "HB"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WriteExcel(tt.args.filename, tt.args.value, tt.args.tableName...); (err != nil) != tt.wantErr {
				t.Errorf("WriteExcel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestReadExcelFile(t *testing.T) {
	type args struct {
		file      File
		tableName []string
	}
	tests := []struct {
		name    string
		args    args
		want    [][]string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test_read_excel_file",
			args: args{
				file: os.NewFile(3, "/tmp/代理人全列表v1.1.xlsx"),
			},
			want:    [][]string{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rows, err := ReadExcelFile(tt.args.file, tt.args.tableName...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadExcelFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for _, row := range rows {
				fmt.Println("row=", row)
			}
		})
	}
}
