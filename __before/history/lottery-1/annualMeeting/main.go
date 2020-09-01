/**
 * 年会抽奖程序
 * 不是线程安全
 * 基础功能：
 * 1 /import 导入参与名单作为抽奖的用户
 * 2 /lucky 从名单中随机抽取用户
 * test.http 用来测试
 */

package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

// 一个用户的列表的切片
var userList []string

// 启动一个iris应用
func newApp() *iris.Application {
	app := iris.New()
	// 把 controller 注册进去
	mvc.New(app.Party("/")).Handle(&lotteryController{})
	return app
}

func main() {
	app := newApp()

	// 给点初始数据
	userList = make([]string, 0)

	// 跑起来
	app.Run(iris.Addr(":8080"))
}

// 因为用的是iris框架，所以需要一个抽奖的控制器
type lotteryController struct {
	// 默认要有一个iris的上下文环境
	Ctx iris.Context
}

// GET http://localhost:8080/
// 看看有多少人参加
func (c *lotteryController) Get() string {
	count := len(userList)
	return fmt.Sprintf("当前总共参与抽奖的用户数: %d\n", count)
}

// POST http://localhost:8080/import
// 导入用户名单
func (c *lotteryController) PostImport() string {
	// 取数据
	strUsers := c.Ctx.FormValue("users")
	// 逗号分隔
	users := strings.Split(strUsers, ",")
	// 导入前统计一下
	count1 := len(userList)
	// 循环
	for _, u := range users {
		// 前后的空格给干掉
		u = strings.TrimSpace(u)
		// 不能是空串
		if len(u) > 0 {
			// 导入用户->追加进去
			userList = append(userList, u)
		}
	}
	// 导入后统计一下
	count2 := len(userList)
	// 看下当前导入多少人
	return fmt.Sprintf("当前总共参与抽奖的用户数: %d，成功导入用户数: %d\n", count2, (count2 - count1))
}

// GET http://localhost:8080/lucky
// 抽中了谁
func (c *lotteryController) GetLucky() string {
	// 统计一下
	count := len(userList)
	// 要有人才能抽
	if count > 1 {
		// 弄个时间戳
		seed := time.Now().UnixNano() // rand内部运算的随机数
		// 生成一个 count 范围内的随机数
		index := rand.New(rand.NewSource(seed)).Int31n(int32(count)) // rand计算得到的随机数
		user := userList[index]                                      // 抽取到一个用户
		userList = append(userList[0:index], userList[index+1:]...)  // 移除这个用户
		return fmt.Sprintf("当前中奖用户: %s, 剩余用户数: %d\n", user, count-1)
	} else if count == 1 { // 就只一个人或只剩一个人
		user := userList[0]
		userList = userList[0:0]
		return fmt.Sprintf("当前中奖用户: %s, 剩余用户数: %d\n", user, count-1)
	} else {
		// 已经没人了
		return fmt.Sprintf("已经没有参与用户，请先通过 /import 导入用户 \n")
	}

}
