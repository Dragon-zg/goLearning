package handle

import (
	"log"
	"net/http"
	"os"
	"strings"
)

type CustomizeError string

func (e CustomizeError) Error() string {
	return string(e)
}

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func ErrWrapper(service appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %v", r)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()

		// 处理静态资源文件
		if strings.Index(request.URL.Path, "favicon.ico") != -1 {
			return
		}

		err := service(writer, request)
		if err != nil {
			log.Printf("Error occurred handling request: %s", err.Error())
			// user error
			if cusErr, ok := err.(CustomizeError); ok {
				http.Error(writer,
					cusErr.Error(),
					http.StatusNetworkAuthenticationRequired)
				return
			}
			// system error
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				panic(err)
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}
