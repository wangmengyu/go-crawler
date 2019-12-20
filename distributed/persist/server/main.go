package main

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"go-craler.com/distributed/config"
	"go-craler.com/distributed/persist"
	"go-craler.com/distributed/rpcsupport"
)

func main() {
	client, err := elastic.NewClient(
		elastic.SetSniff(false), //部署在docker上面是内网，没办法sniff
	)
	if err != nil {
		panic(err)
	}

	err = rpcsupport.ServeRpc(fmt.Sprintf(":%d", config.ItemSaverPort), &persist.ItemSaverService{
		Client: client,
		Index:  config.ElasticIndex,
	})
}
