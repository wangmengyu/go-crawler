package main

import (
	"fmt"
	"go-craler.com/distributed/config"
	"go-craler.com/distributed/persist/client"
	"go-craler.com/engine"
	"go-craler.com/scheduler"
	"go-craler.com/zhenai/parser"
)

func main() {
	//获得城市列表页的HTML内容
	//resp, err := http.Get("http://www.zhenai.com/zhenghun/")
	item, err := client.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan:    item,
	}
	/*
		e.Run(engine.Request{Url: "http://www.zhenai.com/zhenghun/", Parser: parser.ParseCityList})

	*/

	e.Run(engine.Request{Url: "http://www.zhenai.com/zhenghun/shanghai",
		Parser: engine.NewFuncParser(parser.ParseCity, "ParseCity")})

}
