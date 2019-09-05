package main

import (
	"github.com/Kirk-Wang/Hello-Gopher/history/17.7/crawler/engine"
	"github.com/Kirk-Wang/Hello-Gopher/history/17.7/crawler/model"
	"github.com/Kirk-Wang/Hello-Gopher/history/17.7/crawler_distributed/config"
	"github.com/Kirk-Wang/Hello-Gopher/history/17.7/crawler_distributed/rpcsupport"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	const host = ":1234"

	// start ItemSaverServer
	go serveRpc(host, "test1")
	// 让服务器先跑起来
	time.Sleep(time.Second)

	// start ItemSaverClient
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	// Call save
	item := engine.Item{
		Url:  "http://album.zhenai.com/u/1113093123",
		Type: "zhenai",
		Id:   "1113093123",
		Payload: model.Profile{
			Name:       "心想事成",
			Gender:     "女士",
			Age:        26,
			Height:     159,
			Weight:     47,
			Income:     "8千-1.2万",
			Marriage:   "未婚",
			Education:  "硕士",
			Occupation: "金融/银行/保险",
			Hokou:      "重庆",
			Xinzuo:     "射手座",
			House:      "已购房",
			Car:        "未买车",
		},
	}

	result := ""
	err = client.Call(config.ItemSaverRpc, item, &result)

	if err != nil || result != "ok" {
		t.Errorf("result: %s; err: %s", result, err)
	}
}
