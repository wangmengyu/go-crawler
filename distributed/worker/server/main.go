package main

import (
	"fmt"
	"go-craler.com/distributed/config"
	"go-craler.com/distributed/rpcsupport"
	"go-craler.com/distributed/worker"
	"log"
)

func main() {
	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d",
		config.WorkderPort0), worker.CrawlService{}))

}
