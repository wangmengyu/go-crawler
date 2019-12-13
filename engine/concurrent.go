package engine

import (
	"log"
)

/**
  并发版的引擎
*/
type ConcurrentEngine struct {
	Scheduler   Scheduler // 调度器
	WorkerCount int       // 并发数量
}

/**
  调度器，接口类型：
    管理一群调度器回用到的接口的集合
*/
type Scheduler interface {
	Submit(Request)
	WorkerChan() chan Request
	ReadyNotifier // 准备请求
	Run()         //总控
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

/**
  并发引擎的运行方法
*/
func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	//在channel都建立好之后再提交数据
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	//从out chan不断读取数据出来进行打印
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got items: %v", item)
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}

}

/**
创建并发的worker
*/
func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {

	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			result, err := Worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()

}
