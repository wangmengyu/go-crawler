package main

import (
	"flag"
	"fmt"
	"github.com/olivere/elastic/v7"
	"go-craler.com/distributed/config"
	"go-craler.com/distributed/persist"
	"go-craler.com/distributed/rpcsupport"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Printf("must specify a port")
		return
	}

	client, err := elastic.NewClient(
		elastic.SetSniff(false), //部署在docker上面是内网，没办法sniff
	)
	if err != nil {
		panic(err)
	}

	err = rpcsupport.ServeRpc(fmt.Sprintf(":%d", *port), &persist.ItemSaverService{
		Client: client,
		Index:  config.ElasticIndex,
	})
}
