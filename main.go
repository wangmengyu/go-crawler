package main

import (
	"flag"
	"go-craler.com/distributed/config"
	"go-craler.com/distributed/persist/client"
	"go-craler.com/distributed/rpcsupport"
	workerClient "go-craler.com/distributed/worker/client"
	"go-craler.com/engine"
	"go-craler.com/scheduler"
	"go-craler.com/zhenai/parser"
	"log"
	"net/rpc"
	"strings"
)

var (
	itemSaverHost = flag.String("itemsaver_host", "", "itemsaver host")
	workerHosts   = flag.String("worker_hosts", "", "workerer hosts(comma separated)")
)

func main() {
	flag.Parse()
	//获得城市列表页的HTML内容
	//resp, err := http.Get("http://www.zhenai.com/zhenghun/")
	item, err := client.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}
	//创建连接池
	pool := createClientPool(strings.Split(*workerHosts, ","))
	processor, err := workerClient.CreateProcessor(pool)
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      10,
		ItemChan:         item,
		RequestProcessor: processor,
	}
	/*
		e.Run(engine.Request{Url: "http://www.zhenai.com/zhenghun/", Parser: parser.ParseCityList})

	*/

	e.Run(engine.Request{Url: "http://www.zhenai.com/zhenghun/shanghai",
		Parser: engine.NewFuncParser(parser.ParseCity, config.ParseCity)})

}

/**
  创建连接池
*/
func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client

	for _, host := range hosts {
		newClient, err := rpcsupport.NewClient(host)
		if err == nil {
			clients = append(clients, newClient)
			log.Printf("connected to %s", host)
		} else {
			log.Printf("error connection to %s:%v", host, err)
			continue
		}
	}

	//不断的往连接池分发连接，轮流的，用goroutine
	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, c := range clients {
				out <- c
			}
		}
	}()
	return out
}
