package client

import (
	"fmt"
	"go-craler.com/distributed/config"
	"go-craler.com/distributed/worker"
	"go-craler.com/engine"
	"net/rpc"
)

func CreateProcessor(clientChan chan *rpc.Client) (engine.Processor, error) {

	return func(request engine.Request) (result engine.ParseResult, err error) {
		sReq := worker.SerializeRequest(request)
		var sRes worker.ParseResult
		fmt.Println("req:", request)
		fmt.Println("sreq:", sReq)
		client := <-clientChan // 从管道那一个client
		err = client.Call(config.CrawlServiceRpc, sReq, &sRes)
		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(sRes), nil
	}, nil

}
