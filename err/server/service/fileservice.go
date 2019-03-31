package service

import (
	"fmt"
	"goLearning/err/server/handle"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

func GetFileContent(writer http.ResponseWriter, request *http.Request) error {
	//返回自定义错误
	if strings.Index(request.URL.Path, prefix) != 0 {
		return handle.CustomizeError(fmt.Sprintf("path %s must start with %s", request.URL.Path, prefix))
	}
	path := request.URL.Path[len(prefix):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	writer.Write(all)
	return nil
}
