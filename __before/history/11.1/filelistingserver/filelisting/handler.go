package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

// 定义一个自己类型
type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {

	if strings.Index(request.URL.Path, prefix) != 0 {
		// return errors.New("path must start with " + prefix)
		return userError("path must start with " + prefix)
	}

	path := request.URL.Path[len("/list/"):] // /list/fib.txt->fib.txt
	// 打开文件
	file, err := os.Open(path)
	if err != nil {
		// panic(err) -->> 不会崩掉，会在HandleFunc内部受到保护
		// 处理http错误，用户能看到内部的出错信息，不太好
		// http.Error(writer, err.Error(), http.StatusInternalServerError)
		return err
	}
	defer file.Close()
	// 读文件
	all, err := ioutil.ReadAll(file)
	if err != nil {
		// panic(err)
		return err
	}
	// 写会响应
	writer.Write(all)
	return nil // 没错的话，return nil
}
