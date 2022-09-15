package utils

import (
	"fmt"
	"testing"
	"time"
)

// 将time格式化成字符串
func TestTimeToString(t *testing.T) {
	now := time.Now()
	fmt.Printf("YYYY-MM-dd %s\n", now.Format("2006-01-02"))
	fmt.Printf("YYYY-MM-dd HH:mm:ss %s\n", now.Format("2006-01-02 15:04:05"))
	fmt.Printf(Time2Str(now, LocalDateTimePattern))
}

// 将字符串转成time
func TestStringToTime(t *testing.T) {
	str := "2021-01-03 15:23:11"
	// 设置时区
	loc, _ := time.LoadLocation("Asia/Shanghai")
	d, _ := time.ParseInLocation("2006-01-02 15:04:05", str, loc)
	fmt.Printf("time: %v\n", d)

	str2Time, err := Str2Time(str, LocalDateTimePattern)
	if err != nil {
		panic(err)
	}
	fmt.Printf(str2Time.String())
}

// 获取几天前或者几天后
func TestGetDateBefore(t *testing.T) {
	now := time.Now()
	day := 3                          // 获取3天前
	before := now.AddDate(0, 0, -day) // 如果是3天后，则将 - 去掉
	fmt.Printf("time: %v\n", before)
}

// 获取几分钟前或者几分钟后，同样的可以获取几秒前后，几毫秒前后 修改 time.Minute 为 time.Second 或者 time.Hour time.Millisecond 等
func TestGetTimeBefore(t *testing.T) {
	now := time.Now()
	var m time.Duration = -3 // 获取3分钟前
	before := now.Add(time.Minute * m)
	fmt.Printf("time: %v\n", before)
}

func TestLocalDate_UnmarshalJSON(t *testing.T) {
	type fields struct {
		Time time.Time
	}
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test_local_date_unmarshal_json",
			fields: fields{
				Time: time.Now(),
			},
			args: args{
				b: []byte(`"2020-01-01"`),
			},
			wantErr: false,
		},
		{
			name: "test_local_date_unmarshal_json_2",
			fields: fields{
				Time: time.Now(),
			},
			args: args{
				b: []byte(`""`),
			},
			wantErr: false,
		},
		{
			name: "test_local_date_unmarshal_json_3",
			fields: fields{
				Time: time.Now(),
			},
			args: args{
				b: []byte(``),
			},
			wantErr: false,
		},
		{
			name: "test_local_date_unmarshal_json_4",
			fields: fields{
				Time: time.Now(),
			},
			args: args{
				b: []byte(""),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &LocalDate{
				Time: tt.fields.Time,
			}
			if err := d.UnmarshalJSON(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetLocalDateTime(t *testing.T) {
	println(GetLocalDateTimeStr())
	println(GetLocalDateStr())
}

func TestTimeStamp2Str(t *testing.T) {
	type args struct {
		aTime   int64
		pattern string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test_time_stamp_2_str",
			args: args{
				aTime:   1660725142,
				pattern: "2006-01-02 15:04:05",
			},
			want: "2022-08-17 16:32:22",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Timestamp2Str(tt.args.aTime, tt.args.pattern); got != tt.want {
				t.Errorf("TimeStamp2Str() = %v, want %v", got, tt.want)
			}
		})
	}

	t.Run("timestampParse", timestampParse)
}

func timestampParse(t *testing.T) {
	println(time.Now().Unix())
	println(time.Now().UnixMilli())
	println(time.Now().UnixMicro())
	println(time.Now().UnixNano())
}

func TestGetCurrentMonthDateRange(t *testing.T) {
	tests := []struct {
		name          string
		wantDateRange string
	}{
		// TODO: Add test cases.
		{
			name:          "test_get_current_month_date_range",
			wantDateRange: "2022-08-01,2022-08-18",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDateRange := GetCurrentMonthDateRange(); gotDateRange != tt.wantDateRange {
				t.Errorf("GetCurrentMonthDateRange() = %v, want %v", gotDateRange, tt.wantDateRange)
			}
		})
	}
}

func TestGetCurrentMonthDateRange1(t *testing.T) {
	println(GetCurrentMonthDateRange())
}

func TestGetCurrentYearMonth(t *testing.T) {
	fmt.Println(GetCurrentYearMonth())
	fmt.Println(GetLastStatementYearMonth(time.Now()))
}

func TestParseDateYearMonth(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				text: "2020-02-01",
			},
			want:    "202002",
			wantErr: false,
		},
		{
			name: "test2",
			args: args{
				text: "2020-02",
			},
			want:    "202002",
			wantErr: false,
		},
		{
			name: "test3",
			args: args{
				text: "20201023",
			},
			want:    "202010",
			wantErr: false,
		},
		{
			name: "test4",
			args: args{
				text: "20201023452365",
			},
			want:    "202010",
			wantErr: false,
		},
		{
			name: "test4",
			args: args{
				text: "2020",
			},
			want:    "202010",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseDateYearMonth(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDateYearMonth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseDateYearMonth() got = %v, want %v", got, tt.want)
			}
		})
	}
}
