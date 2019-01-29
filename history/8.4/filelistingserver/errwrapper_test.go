package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func errPanic(writer http.ResponseWriter, request *http.Request) error {
	panic(123)
}

func TestErrWrapper(t *testing.T) {
	// 定义数据
	tests := []struct {
		h appHandler // 输入一个函数
		// 期望输出是
		code    int // 404 500
		message string
	}{
		{errPanic, 500, "Internal Server Error"}, // skill: 反正运行完会知道
	}

	for _, tt := range tests {
		// 测试目标函数的行为
		f := errWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(
			http.MethodGet,
			"http://www.lotteryjs.com",
			nil,
		)
		f(response, request)

		b, _ := ioutil.ReadAll(response.Body)
		// 转字符串，去掉换行
		body := strings.Trim(string(b), "\n")
		if response.Code != tt.code || body != tt.message {
			t.Errorf("expect (%d, %s); got (%d, %s)", tt.code, tt.message, response.Code, body)
		}
	}
}
