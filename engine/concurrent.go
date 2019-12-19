package engine

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/go-redis/redis/v7"
	"log"
)

/**
  并发版的引擎
*/
type ConcurrentEngine struct {
	Scheduler   Scheduler // 调度器
	WorkerCount int       // 并发数量
	ItemChan    chan Item //用于save的管道
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

	redisClient := redis.NewClient(&redis.Options{
		Addr: "0.0.0.0:6379",
	})

	//开启并发worker
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	//在channel都建立好之后再提交数据
	for _, r := range seeds {
		if flag, err := isDuplicate(redisClient, r.Url); err == nil && flag == true {
			//log.Printf("重复的url:%s", r.Url)
			continue
		}
		e.Scheduler.Submit(r)
	}

	//从out chan不断读取数据出来进行打印
	for {
		result := <-out

		for _, item := range result.Items {
			go func() {
				e.ItemChan <- item
			}()
		}

		// url 去重 dedup
		for _, request := range result.Requests {
			if flag, err := isDuplicate(redisClient, request.Url); err == nil && flag == true {
				//log.Printf("重复的url:%s", request.Url)
				continue
			}
			e.Scheduler.Submit(request)
		}
	}

}

func isDuplicate(client *redis.Client, url string) (bool, error) {
	data := []byte(url)
	md5Ctx := md5.New()
	md5Ctx.Write(data)
	md5Str := hex.EncodeToString(md5Ctx.Sum(nil))
	val, err := client.Get(md5Str).Result()

	if val == "1" {
		log.Printf("duplicate url:%s,md5:%s", url, md5Str)
		return true, nil
	}

	err = client.Set(md5Str, "1", 0).Err()
	if err != nil {
		log.Printf("redis err2:%v", err)
	}
	return false, err
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
