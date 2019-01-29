package main

import (
	"flag"
	"github.com/Kirk-Wang/Hello-Gopher/17.9/crawler/engine"
	"github.com/Kirk-Wang/Hello-Gopher/17.9/crawler/scheduler"
	"github.com/Kirk-Wang/Hello-Gopher/17.9/crawler/zhenai/parser"
	itemsaver "github.com/Kirk-Wang/Hello-Gopher/17.9/crawler_distributed/persist/client"
	"github.com/Kirk-Wang/Hello-Gopher/17.9/crawler_distributed/rpcsupport"
	worker "github.com/Kirk-Wang/Hello-Gopher/17.9/crawler_distributed/worker/client"
	"log"
	"net/rpc"
	"strings"
)

var (
	itemSaverHost = flag.String("itemsaver_host", "", "itemsaver host")

	workerHosts = flag.String("worker_hosts", "", "worker hosts (comma separated)")
)

func main() {
	flag.Parse()
	itemChan, err := itemsaver.ItemSaver(*itemSaverHost) // 调 rpc, 让远程服务器去 save
	if err != nil {
		panic(err)
	}

	pool := createClientPool(strings.Split(*workerHosts, ","))

	processor := worker.CreateProcessor(pool)

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}

	e.Run(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, "ParseCityList"),
	})
}

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client

	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("Connected to %s", h)
		} else {
			log.Printf("error connecting to %s: %v", h, err)
		}
	}

	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out
}
