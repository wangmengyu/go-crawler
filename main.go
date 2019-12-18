package main

import (
	"go-craler.com/engine"
	"go-craler.com/persist"
	"go-craler.com/scheduler"
	"go-craler.com/zhenai/parser"
)

func main() {
	//获得城市列表页的HTML内容
	//resp, err := http.Get("http://www.zhenai.com/zhenghun/")
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan:    persist.ItemSaver(),
	}
	/*
		e.Run(engine.Request{Url: "http://www.zhenai.com/zhenghun/", ParserFunc: parser.ParseCityList})

	*/

	e.Run(engine.Request{Url: "http://www.zhenai.com/zhenghun/shanghai", ParserFunc: parser.ParseCity})

}
