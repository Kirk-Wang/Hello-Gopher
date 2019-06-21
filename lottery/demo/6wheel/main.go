package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"sync/atomic"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

// 奖品中奖概率
type Prate struct {
	Rate  int    // 万分之N的中奖概率
	Total int    // 总数量限制，0 表示无限数量
	CodeA int    // 中奖概率起始编码（包含）
	CodeB int    // 中奖概率终止编码（包含）
	Left  *int32 // 剩余数
}

// 奖品列表
var prizeList []string
var logger *log.Logger

// var mu sync.Mutex

// 奖品的中奖概率设置，与上面的 prizeList 对应的设置
var rateList []Prate

// 初始化日志
func initLog() {
	f, _ := os.Create("lottery_demo.log")
	// 需要一个日期以及一个毫秒数
	logger = log.New(f, "", log.Ldate|log.Lmicroseconds)
}

func newApp() *iris.Application {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&lotteryController{})
	initLog()
	prizeList = []string{
		"一等奖，火星单程船票",
		"二等奖，凉飕飕南极之旅",
		"三等奖，iPhone一部",
		"", // 没有中奖
	}
	left := int32(1000)
	rateList = []Prate{
		Prate{Rate: 100, Total: 1000, CodeA: 0, CodeB: 9999, Left: &left},
		// Prate{Rate: 1, Total: 1, CodeA: 0, CodeB: 0, Left: 1},
		// Prate{Rate: 2, Total: 2, CodeA: 1, CodeB: 2, Left: 2},
		// Prate{Rate: 5, Total: 10, CodeA: 3, CodeB: 5, Left: 10},
		// Prate{Rate: 100, Total: 0, CodeA: 0, CodeB: 9999, Left: 0},
	}
	return app
}

func main() {
	app := newApp()
	app.Run(iris.Addr(":8080"))
}

// 抽奖的控制器
type lotteryController struct {
	Ctx iris.Context
}

// http://localhost:8080/
func (c *lotteryController) Get() string {
	c.Ctx.Header("Content-Type", "text/html")
	return fmt.Sprintf("大转盘奖品列表：<br/> %s",
		strings.Join(prizeList, "<br />\n"))
}

func (c *lotteryController) GetDebug() string {
	return fmt.Sprintf("获奖概率：%v\n", rateList)
}

func (c *lotteryController) GetPrize() string {
	// mu.Lock()
	// defer mu.Unlock()
	// 第一步，抽奖，根据随机数匹配奖品
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	code := r.Intn(10000)

	var myprize string
	var prizeRate *Prate
	// 从奖品列表匹配是否中奖
	for i, prize := range prizeList {
		// 找对应奖品概率
		rate := &rateList[i]
		if code >= rate.CodeA && code <= rate.CodeB {
			// 满足中奖条件
			myprize = prize
			prizeRate = rate
			break
		}
	}
	if myprize == "" {
		myprize = "很遗憾，再来一次吧"
		return myprize
	}
	// 第二步，中奖了，开始要发奖
	if prizeRate.Total == 0 {
		logger.Printf("%s", myprize)
		// 无限量奖品
		return myprize
	} else if *prizeRate.Left > 0 {
		// 有限量
		// prizeRate.Left--
		left := atomic.AddInt32(prizeRate.Left, -1)
		if left >= 0 {
			logger.Printf("%s,%d", myprize, prizeRate.Left)
			return myprize
		}
	}
	myprize = "很遗憾，再来一次吧"
	return myprize
}
