package utils

import (
	"encoding/binary"
	"errors"
	"fmt"
	logger "github.com/sirupsen/logrus"
	"strings"
	"time"
)

const (
	LocalYearMonth            = "200601"
	LocalYearMonthDay         = "20060102"
	LocalDatePattern          = "2006-01-02"
	LocalDateTimePattern      = "2006-01-02 15:04:05"
	LocalDateMilliTimePattern = "2006-01-02 15:04:05.000"
	LocalDateMicroTimePattern = "2006-01-02 15:04:05.000000"
	LocalDateNanoTimePattern  = "2006-01-02 15:04:05.000000000"
)

type LocalDateTime struct {
	time.Time
}

type LocalDate struct {
	time.Time
}

func GetLocalDateTimeStr() string {
	return time.Now().Format(LocalDateTimePattern)
}

func GetLocalDateStr() string {
	return time.Now().Format(LocalDatePattern)
}

// UnmarshalJSON LocalDateTime implements the json.Unmarshaler interface.
func (dt *LocalDateTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"") //去掉首尾的"
	if s == "" {
		return
	}

	//作一个简单的判断，依据是RFC3339格式的时间类型都带有“T”字符
	if !JudgeRFC3339(s) {
		cst, err := dt.RFC3339ToCSTLayout(s)
		if err != nil {
			fmt.Println(err)
		}
		dt.Time, err = time.Parse(LocalDateTimePattern, cst) //格式化时间格式
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
	} else {
		dt.Time, err = time.Parse(LocalDateTimePattern, s) //格式化时间格式
	}

	logger.Infof("[dt.Time:]={}", dt.Time)
	return
}

// MarshalJSON 自定义序列化,实现MarshalJSON()接口
func (dt LocalDateTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", dt.Time.Format(LocalDateTimePattern))), nil
}

func (dt *LocalDateTime) String() string {
	return Time2Str(dt.Time, LocalDateTimePattern)
}

// RFC3339ToCSTLayout convert rfc3339 value to China standard time layout
func (d *LocalDateTime) RFC3339ToCSTLayout(value string) (string, error) {
	ts, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return "", err
	}
	loc, _ := time.LoadLocation("Asia/Shanghai")
	return ts.In(loc).Format(LocalDateTimePattern), nil
}

// UnmarshalJSON LocalDate //////////////////////////////////////////////
func (d *LocalDate) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"") //去掉首尾的"
	if s == "" {
		return
	}

	//作一个简单的判断，依据是RFC3339格式的时间类型都带有“T”字符
	if !JudgeRFC3339(s) {
		cst, err := d.RFC3339ToCSTLayout(s)
		if err != nil {
			fmt.Println(err)
		}
		d.Time, err = time.Parse(LocalDatePattern, cst) //格式化时间格式
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
	} else {
		d.Time, err = time.Parse(LocalDatePattern, s) //格式化时间格式
	}

	logger.Infof("[dt.Time:]={}", d.Time)
	return
}

func (d LocalDate) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", d.Time.Format(LocalDatePattern))), nil
}

func (d *LocalDate) String() string {
	return Time2Str(d.Time, LocalDatePattern)
}

func (d *LocalDate) GetYearMonth() string {
	return d.Time.Format(LocalYearMonth)
}

func (d *LocalDate) FormatYearMonthDay() string {
	return d.Time.Format(LocalYearMonthDay)
}

// RFC3339ToCSTLayout convert rfc3339 value to China standard time layout
func (d *LocalDate) RFC3339ToCSTLayout(value string) (string, error) {
	ts, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return "", err
	}
	loc, _ := time.LoadLocation("Asia/Shanghai")
	return ts.In(loc).Format(LocalDatePattern), nil
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

// JudgeRFC3339 判断时间类型是否为RFC3339类型
func JudgeRFC3339(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == 'T' {
			return false
		}
	}
	return true
}

// Time2Str TestTimeToString : 将time格式化成字符串
func Time2Str(aTime time.Time, pattern string) string {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	return aTime.In(loc).Format(pattern)
}

// TimeStamp2Str TestTimeToString : 将time格式化成字符串
func Timestamp2Str(aTime int64, pattern string) string {
	return Time2Str(time.Unix(aTime, 0), pattern)
}

// Str2Time TestStringToTime : 将字符串转成time
func Str2Time(timeStr, pattern string) (time.Time, error) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	return time.ParseInLocation(pattern, timeStr, loc)
}

func Str2LocalDate(timeStr, pattern string) (LocalDate, error) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	dt, err := time.ParseInLocation(pattern, timeStr, loc)
	if err != nil {
		return LocalDate{}, err
	}
	return LocalDate{Time: dt}, nil
}

func IsLocalDate(timeStr string) bool {
	date, err := Str2LocalDate(timeStr, LocalDatePattern)
	if err != nil {
		return false
	}
	return date.Time.IsZero() == false
}

func ParseDateRange(dateRange string) (error, LocalDate, LocalDate) {
	split := strings.Split(dateRange, ",")

	if len(split) != 2 {
		return errors.New("dateRange is illegal, eg: '2022-01-01,2023-01-01'"), LocalDate{}, LocalDate{}
	}

	start, err := Str2LocalDate(strings.TrimSpace(split[0]), LocalDatePattern)
	if err != nil {
		return err, LocalDate{}, LocalDate{}
	}

	end, err := Str2LocalDate(strings.TrimSpace(split[1]), LocalDatePattern)
	if err != nil {
		return err, LocalDate{}, LocalDate{}
	}

	return nil, start, end
}

func GetCurrentMonthDateRange() (dateRange string) {
	now := time.Now()
	start := Time2Str(time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local), LocalDatePattern)
	end := Time2Str(time.Date(now.Year(), now.Month()+1, 1, 0, 0, 0, 0, time.Local), LocalDatePattern)
	return fmt.Sprintf("%s,%s", start, end)
}

func GetCurrentYearMonth() string {
	now := time.Now()
	return Time2Str(time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local), LocalYearMonth)
}

func GetLastStatementYearMonth(dateTime time.Time) string {
	return Time2Str(time.Date(dateTime.Year(), dateTime.Month(), 1, 0, 0, 0, 0, time.Local), LocalYearMonth)
}

// ParseDateYearMonth '202201'
func ParseDateYearMonth(text string) (string, error) {
	if len(strings.Split(text, "-")) == 3 {
		parseTime, err := Str2Time(text, LocalDatePattern)
		if err != nil {
			return "", err
		}
		return parseTime.Format(LocalYearMonth), nil
	} else if len(strings.Split(text, "-")) == 2 {
		parseTime, err := Str2Time(text, "2006-01")
		if err != nil {
			return "", err
		}
		return parseTime.Format(LocalYearMonth), nil
	}

	if len(text) > 6 {
		return text[0:6], nil
	} else if len(text) < 6 {
		return "", errors.New(fmt.Sprintf("parse yearMonth=[%s] error.", text))
	}

	return text, nil
}

// Int64ToBytes 将int64转成字节数组
func Int64ToBytes(i int64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

// BytesToInt64 将字节数组转成int64
func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}
