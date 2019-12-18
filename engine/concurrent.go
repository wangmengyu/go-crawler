package engine

import (
	"go-craler.com/model"
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
	//建立输出管道
	out := make(chan ParseResult)
	//启动调度器
	e.Scheduler.Run()

	//开启并发worker
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	//在channel都建立好之后再提交数据
	for _, r := range seeds {
		if isDuplicate(r.Url) {
			//log.Printf("重复的url:%s", r.Url)
			continue
		}
		e.Scheduler.Submit(r)
	}

	//从out chan不断读取数据出来进行打印
	for {
		result := <-out
		profileCount := 0
		for _, item := range result.Items {
			_, ok := item.(model.Profile)
			if ok {
				log.Printf("Got profile: #%d:%v", profileCount, item)
				profileCount++
			}
		}

		// url 去重 dedup
		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				//log.Printf("重复的url:%s", request.Url)
				continue
			}
			e.Scheduler.Submit(request)
		}
	}

}

var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}
	visitedUrls[url] = true
	return false
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
