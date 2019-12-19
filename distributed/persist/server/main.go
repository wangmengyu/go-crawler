package main

import (
	"github.com/olivere/elastic/v7"
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

	err = rpcsupport.ServeRpc(":1234", &persist.ItemSaverService{
		Client: client,
		Index:  "dating_profile",
	})
}
