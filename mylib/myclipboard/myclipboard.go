package myclipboard

import (
	"github.com/phial3/mylib/myclipboard/clipboard"
	"github.com/phial3/mylib/mylog"
)

// Get 读取剪切板中的内容到字符串
func Get() string {
	// 读取剪切板中的内容到字符串
	content, err := clipboard.ReadAll()
	if err != nil {
		mylog.Error(err)
		return ""
	}
	return content
}

// Set 复制内容到剪切板
func Set(cb string) {
	clipboard.WriteAll(cb)
}
