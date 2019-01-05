/**
 * 大转盘程序
 * curl http://localhost:8080/
 * curl http://localhost:8080/debug
 * curl http://localhost:8080/prize
 * 固定几个奖品，不同的中奖概率或者总数量限制
 * 每一次转动抽奖，后端计算出这次抽奖的中奖情况，并返回对应的奖品信息
 *
 * 线程不安全，因为获奖概率低，并发更新库存的冲突很少能出现，不容易发现线程安全性问题
 * 压力测试：
 * wrk -t10 -c100 -d5 "http://localhost:8080/prize"
 */
package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"math/rand"
	"strings"
	"sync"
	"time"
)

// 奖品中奖概率
type Prate struct {
	Rate  int // 万分之N的中奖概率
	Total int // 总数量限制，0 表示无限数量
	CodeA int // 中奖概率起始编码（包含）
	CodeB int // 中奖概率终止编码（包含）
	Left  int // 剩余数
}

// 奖品列表
var prizeList []string = []string{
	"一等奖，火星单程船票",
	"二等奖，凉飕飕南极之旅",
	"三等奖，iPhone一部",
	"", // 没有中奖
}

var mu sync.Mutex = sync.Mutex{}

// 奖品的中奖概率设置，与上面的 prizeList 对应的设置
var rateList []Prate = []Prate{
	Prate{100, 1000, 0, 9999, 1000},
	// Prate{1, 1, 0, 0, 1},
	// Prate{2, 2, 1, 2, 2},
	// Prate{5, 10, 3, 5, 10},
	// Prate{100, 0, 0, 9999, 0},
}

func newApp() *iris.Application {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&lotteryController{})
	return app
}

func main() {
	app := newApp()
	// http://localhost:8080
	app.Run(iris.Addr(":8080"))
}

// 抽奖的控制器
type lotteryController struct {
	Ctx iris.Context
}

// GET http://localhost:8080/
func (c *lotteryController) Get() string {
	c.Ctx.Header("Content-Type", "text/html")
	return fmt.Sprintf("大转盘奖品列表：<br/> %s", strings.Join(prizeList, "<br/>\n"))
}

// GET http://localhost:8080/prize
func (c *lotteryController) GetPrize() string {
	c.Ctx.Header("Content-Type", "text/html")

	// 第一步，抽奖，根据随机数匹配奖品
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	// 得到个人的抽奖编码
	code := r.Intn(10000) // 10000 以内的一个数
	//fmt.Println("GetPrize code=", code)
	var myprize string
	var prizeRate *Prate
	// 从奖品列表中匹配，是否中奖
	for i, prize := range prizeList {
		rate := &rateList[i]
		// 抽奖编码是否在区间内
		if code >= rate.CodeA && code <= rate.CodeB {
			// 满足中奖条件
			myprize = prize
			prizeRate = rate
			break
		}
	}
	if myprize == "" {
		// 没有中奖
		myprize = "很遗憾，再来一次"
		return myprize
	}
	// 第二步，发奖，是否可以发奖
	if prizeRate.Total == 0 {
		// 无限奖品
		fmt.Println("中奖： ", myprize)
		return myprize
	} else if prizeRate.Left > 0 {
		mu.Lock()
		// 还有剩余奖品
		prizeRate.Left -= 1
		mu.Unlock()
		fmt.Println("中奖： ", myprize)
		return myprize
	} else {
		// 有限且没有剩余奖品，无法发奖
		myprize = "很遗憾，再来一次"
		return myprize
	}
}

// GET http://localhost:8080/debug
func (c *lotteryController) GetDebug() string {
	c.Ctx.Header("Content-Type", "text/html")
	return fmt.Sprintf("获奖概率： %v", rateList)
}
