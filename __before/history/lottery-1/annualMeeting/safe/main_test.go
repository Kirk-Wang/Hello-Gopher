/**
 * 线程是否安全的测试
 * slice切片在并发请求时不会出现异常，但也是线程不安全的，数据更新会有问题
 * go test -v
 */
package main

import (
	"fmt"
	"github.com/kataras/iris/httptest"
	"sync"
	"testing"
)

func TestMVC(t *testing.T) {
	// 利用 iris httptest
	e := httptest.New(t, newApp())

	// 来保证所有的协程全部执行完
	var wg sync.WaitGroup
	// 期望 StatusOK
	e.GET("/").Expect().Status(httptest.StatusOK).
		Body().Equal("当前总共参与抽奖的用户数: 0\n")

	// 启动100个协程并发来执行用户导入操作
	// 如果是线程安全的时候，预期倒入成功100个用户
	for i := 0; i < 100; i++ {
		// 每次并发，都向 WaitGroup 添加一个 task
		wg.Add(1)
		// 跑协程
		go func(i int) {
			// 协程跑完了，注意执行一个 Done 操作
			defer wg.Done()
			// 导名单
			e.POST("/import").WithFormField("users", fmt.Sprintf("test_u%d", i)).Expect().Status(httptest.StatusOK)
		}(i)
	}
	// 等待上面所有协程都 Done
	wg.Wait()

	// 期望有100个用户
	e.GET("/").Expect().Status(httptest.StatusOK).
		Body().Equal("当前总共参与抽奖的用户数: 100\n")
	// 期望抽奖成功
	e.GET("/lucky").Expect().Status(httptest.StatusOK)
	// 期望还有 99 个用户
	e.GET("/").Expect().Status(httptest.StatusOK).
		Body().Equal("当前总共参与抽奖的用户数: 99\n")
}
