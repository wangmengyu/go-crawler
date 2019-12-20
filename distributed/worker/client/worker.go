package client

import (
	"fmt"
	"go-craler.com/distributed/config"
	"go-craler.com/distributed/rpcsupport"
	"go-craler.com/distributed/worker"
	"go-craler.com/engine"
)

func CreateProcessor() (engine.Processor, error) {
	client, err := rpcsupport.NewClient(fmt.Sprintf(":%d", config.WorkderPort0))
	if err != nil {
		return nil, err
	}
	return func(request engine.Request) (result engine.ParseResult, err error) {
		sReq := worker.SerializeRequest(request)
		var sRes worker.ParseResult
		fmt.Println("req:", request)
		fmt.Println("sreq:", sReq)
		err = client.Call(config.CrawlServiceRpc, sReq, &sRes)
		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(sRes), nil
	}, nil

}
