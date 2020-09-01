/**
 * 支付宝五福
 * 五福的概率来自识别后的参数(AI图片识别MaBaBa)
 * 基础功能：
 * /lucky 只有一个抽奖的接口，奖品信息都是预先配置好的
 * 测试方法：
 * curl "http://localhost:8080/?rate=4,3,2,1,0"
 * curl "http://localhost:8080/lucky?uid=1&rate=4,3,2,1,0"
 * 压力测试：（这里不存在线程安全问题）
 * wrk -t10 -c10 -d 10 "http://localhost:8080/lucky?uid=1&rate=4,3,2,1,0"
 */

package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

// 最大号码
const rateMax = 10

// 奖品信息
type gift struct {
	id      int    // 奖品ID
	name    string // 奖品名称
	pic     string // 照片链接
	link    string // 链接
	inuse   bool   // 是否使用中
	rate    int    // 中奖概率，十分之N,0-9
	rateMin int    // 大于等于，中奖的最小号码,0-10
	rateMax int    // 小于，中奖的最大号码,0-10
}

// 文件日志
var logger *log.Logger

func main() {
	app := newApp()

	// http://localhost:8080/
	app.Run(iris.Addr(":8080"))
}

// 初始化奖品列表信息（管理后台来维护）
func newGift() *[5]gift {
	giftlist := new([5]gift) // 这里用个数组
	// 1 实物大奖
	g1 := gift{
		id:      1,
		name:    "富强福",
		pic:     "富强福.jpg",
		link:    "",
		inuse:   true,
		rate:    4,
		rateMin: 0,
		rateMax: 0,
	}
	giftlist[0] = g1
	// 2 实物小奖
	g2 := gift{
		id:      2,
		name:    "和谐福",
		pic:     "和谐福.jpg",
		link:    "",
		inuse:   true,
		rate:    3,
		rateMin: 0,
		rateMax: 0,
	}
	giftlist[1] = g2
	// 3 虚拟券，相同的编码
	g3 := gift{
		id:      3,
		name:    "友善福",
		pic:     "友善福.jpg",
		link:    "",
		inuse:   true,
		rate:    2,
		rateMin: 0,
		rateMax: 0,
	}
	giftlist[2] = g3
	// 4 虚拟券，不相同的编码
	g4 := gift{
		id:      4,
		name:    "爱国福",
		pic:     "爱国福.jpg",
		link:    "",
		inuse:   true,
		rate:    1,
		rateMin: 0,
		rateMax: 0,
	}
	giftlist[3] = g4
	// 5 虚拟币
	g5 := gift{
		id:      5,
		name:    "敬业福",
		pic:     "敬业福.jpg",
		link:    "",
		inuse:   true,
		rate:    0,
		rateMin: 0,
		rateMax: 0,
	}
	giftlist[4] = g5
	return giftlist
}

// 根据概率，计算好的奖品信息列表
func giftRate(rate string) *[5]gift {
	giftlist := newGift()
	rates := strings.Split(rate, ",")
	ratesLen := len(rates)
	// 整理奖品数据，把rateMin,rateMax根据rate进行编排
	rateStart := 0
	for i, data := range giftlist {
		if !data.inuse {
			continue
		}
		grate := 0
		if i < ratesLen { // 避免数组越界
			grate, _ = strconv.Atoi(rates[i])
		}
		giftlist[i].rate = grate
		giftlist[i].rateMin = rateStart
		giftlist[i].rateMax = rateStart + grate
		if giftlist[i].rateMax >= rateMax {
			// 号码达到最大值，分配的范围重头再来
			giftlist[i].rateMax = rateMax
			rateStart = 0
		} else {
			rateStart += grate
		}
	}
	fmt.Printf("giftlist=%v\n", giftlist)
	return giftlist
}

// 初始化日志信息
func initLog() {
	f, _ := os.Create("./lottery_demo.log")
	logger = log.New(f, "", log.Ldate|log.Lmicroseconds)
}

func newApp() *iris.Application {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&lotteryController{})
	// 初始化日志信息
	initLog()
	return app
}

// 抽奖的控制器
type lotteryController struct {
	Ctx iris.Context
}

// GET http://localhost:8080/?rate=4,3,2,1,0
func (c *lotteryController) Get() string {
	rate := c.Ctx.URLParamDefault("rate", "4,3,2,1,0") // 这里给个默认值
	giftlist := giftRate(rate)
	return fmt.Sprintf("%v\n", giftlist)
}

// GET http://localhost:8080/lucky?uid=1&rate=4,3,2,1,0
func (c *lotteryController) GetLucky() map[string]interface{} {
	uid, _ := c.Ctx.URLParamInt("uid")
	rate := c.Ctx.URLParamDefault("rate", "4,3,2,1,0")
	code := luckyCode()
	ok := false
	result := make(map[string]interface{})
	result["success"] = ok
	giftlist := giftRate(rate)
	for _, data := range giftlist {
		if !data.inuse {
			continue
		}
		if data.rateMin <= int(code) && data.rateMax >= int(code) {
			// 中奖了，抽奖编码在奖品中奖编码范围内
			ok = true
			sendData := data.pic

			if ok {
				// 中奖后，成功得到奖品（发奖成功）
				// 生成中奖纪录
				saveLuckyData(uid, code, data.id, data.name, data.link, sendData)
				result["success"] = ok
				result["uid"] = uid
				result["id"] = data.id
				result["name"] = data.name
				result["link"] = data.link
				result["data"] = sendData
				break
			}
		}
	}

	return result
}

// 抽奖编码
func luckyCode() int32 {
	seed := time.Now().UnixNano()                                 // rand内部运算的随机数
	code := rand.New(rand.NewSource(seed)).Int31n(int32(rateMax)) // rand计算得到的随机数
	return code
}

// 记录用户的获奖记录
func saveLuckyData(uid int, code int32, id int, name, link, sendData string) {
	logger.Printf("lucky, uid=%d, code=%d, gift=%d, name=%s, link=%s, data=%s ", uid, code, id, name, link, sendData)
}
