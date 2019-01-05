/**
 * 微信摇一摇
 * 增加互斥锁，保证并发更新数据的安全
 * 基础功能：
 * /lucky 只有一个抽奖的接口，奖品信息都是预先配置好的
 * 测试方法：
 * curl http://localhost:8080/
 * curl http://localhost:8080/lucky
 * 压力测试：（线程不安全的时候，总的中奖纪录会超过总的奖品数）
 * wrk -t10 -c10 -d5 http://localhost:8080/lucky
 */

package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"sync"
)

// 奖品类型，枚举 iota 从0开始自增
const (
	giftTypeCoin      = iota // 虚拟币
	giftTypeCoupon           // 优惠券，不相同的编码
	giftTypeCouponFix        // 优惠券，相同的编码
	giftTypeRealSmall        // 实物小奖
	giftTypeRealLarge        // 实物大奖
)

// 最大号码
const rateMax = 10000

// 奖品信息
type gift struct {
	id       int      // 奖品ID
	name     string   // 奖品名称
	pic      string   // 照片链接
	link     string   // 链接
	gtype    int      // 奖品类型
	data     string   // 奖品的数据（特定的配置信息，如：虚拟币面值，固定优惠券的编码）
	datalist []string // 奖品数据集合（特定的配置信息，如：不同的优惠券的编码）
	total    int      // 总数，0 不限量
	left     int      // 剩余数
	inuse    bool     // 是否使用中
	rate     int      // 中奖概率，万分之N,0-10000
	rateMin  int      // 大于等于，中奖的最小号码,0-10000
	rateMax  int      // 小于，中奖的最大号码,0-10000
}

// 文件日志
var logger *log.Logger

// 奖品列表
var giftlist []*gift
var mu sync.Mutex

func main() {
	app := newApp()

	// http://localhost:8080
	app.Run(iris.Addr(":8080"))
}

// 初始化奖品列表信息（管理后台来维护）
func initGift() {
	giftlist = make([]*gift, 5)
	// 1 实物大奖
	g1 := gift{
		id:      1,
		name:    "手机N7",
		pic:     "",
		link:    "",
		gtype:   giftTypeRealLarge,
		data:    "",
		total:   1000,
		left:    1000,
		inuse:   true,
		rate:    10000,
		rateMin: 0,
		rateMax: 0,
	}
	giftlist[0] = &g1
	// 2 实物小奖
	g2 := gift{
		id:      2,
		name:    "安全充电 黑色",
		pic:     "",
		link:    "",
		gtype:   giftTypeRealSmall,
		data:    "",
		total:   5,
		left:    5,
		inuse:   true,
		rate:    100,
		rateMin: 0,
		rateMax: 0,
	}
	giftlist[1] = &g2
	// 3 虚拟券，相同的编码
	g3 := gift{
		id:      3,
		name:    "商城满2000元减50元优惠券",
		pic:     "",
		link:    "",
		gtype:   giftTypeCouponFix,
		data:    "mall-coupon-2018",
		total:   5,
		left:    5,
		rate:    5000,
		inuse:   true,
		rateMin: 0,
		rateMax: 0,
	}
	giftlist[2] = &g3
	// 4 虚拟券，不相同的编码
	g4 := gift{
		id:       4,
		name:     "商城无门槛直降50元优惠券",
		pic:      "",
		link:     "",
		gtype:    giftTypeCoupon,
		data:     "",
		datalist: []string{"c01", "c02", "c03", "c04", "c05"},
		total:    5,
		left:     5,
		inuse:    true,
		rate:     2000,
		rateMin:  0,
		rateMax:  0,
	}
	giftlist[3] = &g4
	// 5 虚拟币
	g5 := gift{
		id:      5,
		name:    "社区10个金币",
		pic:     "",
		link:    "",
		gtype:   giftTypeCoin,
		data:    "10",
		total:   5,
		left:    5,
		inuse:   true,
		rate:    5000,
		rateMin: 0,
		rateMax: 0,
	}
	giftlist[4] = &g5

	// 整理奖品数据，把rateMin,rateMax根据rate进行编排
	rateStart := 0
	for _, data := range giftlist {
		if !data.inuse {
			continue
		}
		data.rateMin = rateStart
		data.rateMax = data.rateMin + data.rate
		if data.rateMax >= rateMax {
			// 号码达到最大值，分配的范围重头再来
			data.rateMax = rateMax
			rateStart = 0
		} else {
			rateStart += data.rate
		}
	}
	fmt.Printf("giftlist=%v\n", giftlist)
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
	// 初始化奖品信息
	initGift()
	return app
}

// 抽奖的控制器
type lotteryController struct {
	Ctx iris.Context
}

// GET http://localhost:8080/
func (c *lotteryController) Get() string {
	count := 0
	total := 0
	for _, data := range giftlist {
		if data.inuse && (data.total == 0 ||
			(data.total > 0 && data.left > 0)) {
			count++
			total += data.left
		}
	}
	return fmt.Sprintf("当前有效奖品种类数量: %d，限量奖品总数量=%d\n", count, total)
}

// GET http://localhost:8080/lucky
func (c *lotteryController) GetLucky() map[string]interface{} {
	mu.Lock()
	defer mu.Unlock()

	code := luckyCode()
	ok := false
	result := make(map[string]interface{})
	result["success"] = ok
	for _, data := range giftlist {
		if !data.inuse || (data.total > 0 && data.left <= 0) {
			continue
		}
		if data.rateMin <= int(code) && data.rateMax > int(code) {
			// 中奖了，抽奖编码在奖品中奖编码范围内
			sendData := ""
			switch data.gtype {
			case giftTypeCoin:
				ok, sendData = sendCoin(data)
			case giftTypeCoupon:
				ok, sendData = sendCoupon(data)
			case giftTypeCouponFix:
				ok, sendData = sendCouponFix(data)
			case giftTypeRealSmall:
				ok, sendData = sendRealSmall(data)
			case giftTypeRealLarge:
				ok, sendData = sendRealLarge(data)
			}
			if ok {
				// 中奖后，成功得到奖品（发奖成功）
				// 生成中奖纪录
				saveLuckyData(code, data.id, data.name, data.link, sendData, data.left)
				result["success"] = ok
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

// 发奖，虚拟币
func sendCoin(data *gift) (bool, string) {
	if data.total == 0 {
		// 数量无限
		return true, data.data
	} else if data.left > 0 {
		// 还有剩余
		data.left = data.left - 1
		return true, data.data
	} else {
		return false, "奖品已发完"
	}
}

// 发奖，优惠券（不同值）
func sendCoupon(data *gift) (bool, string) {
	if data.left > 0 {
		// 还有剩余的奖品
		left := data.left - 1
		data.left = left
		return true, data.datalist[left]
	} else {
		return false, "奖品已发完"
	}
}

// 发奖，优惠券（固定值）
func sendCouponFix(data *gift) (bool, string) {
	if data.total == 0 {
		// 数量无限
		return true, data.data
	} else if data.left > 0 {
		data.left = data.left - 1
		return true, data.data
	} else {
		return false, "奖品已发完"
	}
}

// 发奖，实物小
func sendRealSmall(data *gift) (bool, string) {
	if data.total == 0 {
		// 数量无限
		return true, data.data
	} else if data.left > 0 {
		data.left = data.left - 1
		return true, data.data
	} else {
		return false, "奖品已发完"
	}
}

// 发奖，实物大
func sendRealLarge(data *gift) (bool, string) {
	if data.total == 0 {
		// 数量无限
		return true, data.data
	} else if data.left > 0 {
		data.left--
		return true, data.data
	} else {
		return false, "奖品已发完"
	}
}

// 记录用户的获奖记录
func saveLuckyData(code int32, id int, name, link, sendData string, left int) {
	logger.Printf("lucky, code=%d, gift=%d, name=%s, link=%s, data=%s, left=%d ", code, id, name, link, sendData, left)
}
