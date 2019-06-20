package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

// 红包列表
var packageList map[uint32][]uint

type lotteryContrller struct {
	Ctx iris.Context
}

func newApp() *iris.Application {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&lotteryContrller{})
	return app
}

func main() {
	app := newApp()
	app.Run(iris.Addr(":8080"))
}

// 返回全部红包地址
// http://localhost:8080/
func (c *lotteryContrller) Get() map[uint32][2]int {
	rs := make(map[uint32][2]int)
	for id, list := range packageList {
		var money int
		for _, v := range list {
			money += int(v)
		}
		rs[id] = [2]int{len(list), money}
	}
	return rs
}

// 发红包
// http://localhost:8080/set?uid=1&money=100&num=100
func (c *lotteryContrller) GetSet() string {
	uid, errUid := c.Ctx.URLParamInt("uid")
	money, errMoney := c.Ctx.URLParamFloat64("money")
	num, errNum := c.Ctx.URLParamInt("num")

	if errUid != nil || errMoney != nil || errNum != nil {
		return fmt.Sprintf("参数格式异常，errUid=%d, errMoney=%d, errNum=%d\n",
			errUid, errMoney, errNum)
	}
	// 精确到分
	moneyTotal := int(money * 100)
	if uid < 1 || moneyTotal < num || num < 1 {
		return fmt.Sprintf("参数数值异常，uid=%d, money=%d, num=%d\n", uid, money, num)
	}
	// 金额分配算法
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rMax := 0.55 // 最大红包占比(随机分配的最大值)
	list := make([]uint, num)
	leftMoney := moneyTotal // 剩余的金额
	leftNum := num          // 剩余的红包数
	// 大循环开始，分配金额给到每一个红包
	for leftNum > 0 {
		if leftNum == 1 {
			// 最后一个红包，剩余的全部金额都给它
			list[num-1] = uint(leftMoney)
			break
		}
		if leftMoney == leftNum {
			// 最小单位
			for i := num - leftNum; i < num; i++ {
				list[i] = 1
			}
			break
		}
		// (leftMoney-leftNum): 给剩下的人至少一分钱
		rMoney := int(float64(leftMoney-leftNum) * rMax)
		m := r.Intn(rMoney) // 生成 0-rMoney 的一个值
		if m < 1 {
			m = 1 // 把 0 去掉
		}
		list[num-leftNum] = uint(m)
		leftMoney -= m
		leftNum--
	}
	// 红包的唯一ID
	id := r.Uint32()
	packageList[id] = list
	// 返回抢红包的URL
	return fmt.Sprintf("/get?id=%d&uid=%d&num=%d", id, uid, num)
}
