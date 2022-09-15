package main

import (
	"embed"
	"fmt"
	"io/fs"
)

//go:embed "data/*"
var files embed.FS

func main() {
	templates, _ := fs.ReadDir(files, "data")

	//打印出文件名称
	for _, template := range templates {
		fmt.Printf("%q\n", template.Name())
	}
}
