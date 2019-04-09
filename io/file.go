package main

import (
	"path/filepath"
	"fmt"
	"os"
	"log"
	"io/ioutil"
)

const ioFilePath = "io/file.go"

func main() {
	filePath := getFileAbsPath(ioFilePath)
	file, err := os.Open(getFileAbsPath(filePath))
	if err != nil {
		log.Println("file not exists:", filePath)
		return
	}
	defer file.Close()

	contents, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("read file error!")
		return
	}

	fmt.Println("文件内容如下: ", string(contents))
}

func  getFileAbsPath(filePath string) string {
	path, err := filepath.Abs(filePath)
	if err != nil {
		panic(err)
	}
	return path
}