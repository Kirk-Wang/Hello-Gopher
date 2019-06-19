package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/kataras/iris/mvc"

	"github.com/kataras/iris"
)

type gift struct {
	id      int    // 奖品 ID
	name    string // 奖品名称
	pic     string // 奖品的图片
	link    string // 奖品的链接
	inuse   bool   // 是否使用中
	rate    int    // 中奖概率，万分之N，0-9999
	rateMin int    // 大于等于最小中奖编码
	rateMax int    // 小于中奖编码
}

// 最大的中奖号码
const rateMax = 10

var logger *log.Logger

type lotteryController struct {
	Ctx iris.Context
}

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

	return app
}

func main() {
	app := newApp()
	app.Run(iris.Addr(":8080"))
}

func newGift() *[5]gift {
	giftList := new([5]gift)
	g1 := gift{
		id:      1,
		name:    "富强福",
		pic:     "富强福.jpg",
		link:    "",
		inuse:   true,
		rate:    0,
		rateMin: 0,
		rateMax: 0,
	}
	giftList[0] = g1
	g2 := gift{
		id:      2,
		name:    "和谐福",
		pic:     "和谐福.jpg",
		link:    "",
		inuse:   true,
		rate:    0,
		rateMin: 0,
		rateMax: 0,
	}
	giftList[1] = g2
	g3 := gift{
		id:      3,
		name:    "友善福",
		pic:     "友善福.jpg",
		link:    "",
		inuse:   true,
		rate:    0,
		rateMin: 0,
		rateMax: 0,
	}
	giftList[2] = g3
	g4 := gift{
		id:      4,
		name:    "爱国福",
		pic:     "爱国福.jpg",
		link:    "",
		inuse:   true,
		rate:    0,
		rateMin: 0,
		rateMax: 0,
	}
	giftList[3] = g4
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
	giftList[4] = g5

	return giftList
}

// 动态设置概率
func giftRate(rate string) *[5]gift {
	giftList := newGift()
	rates := strings.Split(rate, ",")
	ratesLen := len(rates)
	// 数据整理，中奖区间数据
	rateStart := 0
	for i, data := range giftList {
		if !data.inuse {
			continue
		}
		grate := 0
		if i < ratesLen {
			grate, _ = strconv.Atoi(rates[i])
		}
		fmt.Printf("%d\n", grate)
		giftList[i].rate = grate
		giftList[i].rateMin = rateStart
		giftList[i].rateMax = rateStart + grate

		if giftList[i].rateMax > rateMax {
			giftList[i].rateMax = rateMax
			rateStart = 0
		} else {
			rateStart += grate
		}
	}
	fmt.Printf("giftList=%v\n", giftList)
	return giftList
}

// GET http://localhost:8080/?rate=4,3,2,1,0
func (c *lotteryController) Get() string {
	rate := c.Ctx.URLParamDefault("rate", "4,3,2,1,0")
	giftList := giftRate(rate)
	return fmt.Sprintf("%v\n", giftList)
}

// 抽奖 GET http://localhost:8080/lucky?uid=1&rate=4,3,2,1,0
func (c *lotteryController) GetLucky() map[string]interface{} {
	uid, _ := c.Ctx.URLParamInt("uid")
	rate := c.Ctx.URLParamDefault("rate", "4,3,2,1,0")
	code := int(luckyCode())
	result := make(map[string]interface{})
	result["success"] = false
	giftList := giftRate(rate)

	for _, data := range giftList {
		// 没被使用
		if !data.inuse {
			continue
		}
		if code >= data.rateMin && code < data.rateMax {
			// 中奖了,抽奖编码在奖品编码范围内
			// 开始发奖
			sendData := data.pic
			// 中奖后，成功得到奖品
			// 生成中奖纪录
			saveLuckyData(code, data.id, data.name, data.link, sendData)
			result["success"] = true
			result["uid"] = uid
			result["id"] = data.id
			result["name"] = data.name
			result["link"] = data.link
			result["data"] = sendData
			break
		}
	}

	return result
}

func luckyCode() int32 {
	seed := time.Now().UnixNano()
	code := rand.New(rand.NewSource(seed)).Int31n(rateMax)
	return code
}

// 纪录用户的获奖信息
func saveLuckyData(code, id int, name, link, sendData string) {
	logger.Printf("lucky, code=%d, gift=%d, name=%s, link=%s, data=%s \n",
		code, id, name, link, sendData)
}
