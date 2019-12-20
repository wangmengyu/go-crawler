package engine

import (
	"go-craler.com/fetcher"
	"log"
)

func worker(r Request) (ParseResult, error) {
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
