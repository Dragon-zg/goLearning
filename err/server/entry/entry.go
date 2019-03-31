package main

import (
	"goLearning/err/server/config"
	"goLearning/err/server/handle"
	"goLearning/err/server/service"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handle.ErrWrapper(service.GetFileContent))

	err := http.ListenAndServe(config.ServerPort, nil)
	if err != nil {
		log.Fatalln("服务启动失败!", err.Error())
	}
}
