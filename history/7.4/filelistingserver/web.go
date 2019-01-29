package main

import (
	"github.com/Kirk-Wang/Hello-Gopher/7.4/filelistingserver/filelisting"
	"log"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := handler(writer, request)
		if err != nil {
			// log.W("Error handling request: %s", err.Error())
			log.Printf("Error handling request: %s", err.Error())
			// log.Warn("Error handling request: %s", err.Error())
			// Warning := log.New(os.Stdout, "Warning:", log.Ldate|log.Ltime|log.Lshortfile)
			// Warning.Printf("Error handling request: %s", err.Error())
			code := http.StatusOK
			switch {
			// 错误类型的判断
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(
				writer,                // 向谁汇报这个错误
				http.StatusText(code), // 不想直接 err.Error()->暴露了内部的错
				code)                  // 404 code 本身
		}
	}
}

func main() {
	http.HandleFunc("/list/", errWrapper(filelisting.HandleFileList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
