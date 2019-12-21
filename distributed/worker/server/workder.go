package main

import (
	"flag"
	"fmt"
	"go-craler.com/distributed/rpcsupport"
	"go-craler.com/distributed/worker"
	"log"
)

//命令行参数
var port = flag.Int("port", 0, "the port for me to listen on ")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d",
		*port), worker.CrawlService{}))

}
