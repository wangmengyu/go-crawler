package engine

import (
	"go-craler.com/fetcher"
	"log"
)

/**
  简单调度器
*/
type SimpleEngine struct {
}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request

	//将所有seed放入处理队列
	for _, r := range seeds {
		requests = append(requests, r)
	}

	//处理队列中所有数据
	for len(requests) > 0 {
		//取得一个请求
		r := requests[0]
		requests = requests[1:]
		parseResult, err := Worker(r)
		if err != nil {
			continue
		}

		//将返回中的URL送入requests队列中继续消耗
		requests = append(requests, parseResult.Requests...)

		//打印parseResult
		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}

	}

}

func Worker(r Request) (ParseResult, error) {
	//获取网页内容
	//log.Printf("Fetching %s", r.Url)

	bodys, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("fetcher error url %s %v", r.Url, err)
		return ParseResult{}, err
	}

	//将bodys内容进行解析
	parseResult := r.Parser.Parse(bodys, r.Url)
	return parseResult, nil
}
