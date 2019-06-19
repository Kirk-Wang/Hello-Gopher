package main

import (
	"log"
	"os"

	"github.com/kataras/iris/mvc"

	"github.com/kataras/iris"
)

// 奖品类型，枚举值 iota 从 0 开始
const (
	giftTypeCoin      = iota // 虚拟币
	giftTypeCoupon           // 不同券
	giftTypeCouponFix        // 相同券
	giftTypeRealSamll        // 实物小奖
	giftTypeRealLarge        // 实物大奖
)

type gift struct {
	id       int      // 奖品 ID
	name     string   // 奖品名称
	pic      string   // 奖品的图片
	link     string   // 奖品的链接
	gtype    int      // 奖品类型
	data     string   // 奖品数据（特定的配置信息）
	datalist []string // 奖品数据集合（不同的优惠券的编码）
	total    int      // 总数，0 不限量
	left     int      // 剩余数量
	inuse    bool     // 是否使用中
	rate     int      // 中奖概率，万分之N，0-9999
	rateMin  int      // 大于等于最小中奖编码
	rateMax  int      // 小于中奖编码
}

// 最大的中奖号码
const rateMax = 10000

var logger *log.Logger

// 奖品列表
var giftList []*gift

type lotteryController struct {
	Ctx iris.Context
}

// 初始化日志
func initLog() {
	f, _ := os.Create("/var/log/lottery_demo.log")
	// 需要一个日期以及一个毫秒数
	logger = log.New(f, "", log.Ldate|log.Lmicroseconds)
}

// 初始化奖品列表
func initGift() {
	giftList = make([]*gift, 5)
	g1 := gift{
		id:       1,
		name:     "手机大奖",
		pic:      "",
		link:     "",
		gtype:    giftTypeRealLarge,
		data:     "",
		datalist: nil,
		total:    1000,
		left:     1000,
		inuse:    true,
		rate:     10000,
		rateMin:  0,
		rateMax:  0,
	}
	giftList[0] = &g1
	g2 := gift{
		id:       2,
		name:     "充电器",
		pic:      "",
		link:     "",
		gtype:    giftTypeRealSamll,
		data:     "",
		datalist: nil,
		total:    5,
		left:     5,
		inuse:    true,
		rate:     100,
		rateMin:  0,
		rateMax:  0,
	}
	giftList[1] = &g2
	g3 := gift{
		id:       3,
		name:     "优惠券满200减50",
		pic:      "",
		link:     "",
		gtype:    giftTypeCouponFix,
		data:     "mall-coupon-2018",
		datalist: nil,
		total:    5,
		left:     5,
		inuse:    true,
		rate:     5000,
		rateMin:  0,
		rateMax:  0,
	}
	giftList[2] = &g3
	g4 := gift{
		id:       4,
		name:     "直降优惠券50元",
		pic:      "",
		link:     "",
		gtype:    giftTypeCoupon,
		data:     "",
		datalist: []string{"c01", "c02", "c03", "c04", "c05"},
		total:    5,
		left:     5,
		inuse:    true,
		rate:     5000,
		rateMin:  0,
		rateMax:  0,
	}
	giftList[3] = &g4
}

func newApp() *iris.Application {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&lotteryController{})

	initLog()

	return app
}

func main() {

}
