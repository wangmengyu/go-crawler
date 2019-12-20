package client

import (
	"go-craler.com/distributed/config"
	"go-craler.com/distributed/rpcsupport"
	"go-craler.com/engine"
	"log"
)

func ItemSaver(host string) (chan engine.Item, error) {
	out := make(chan engine.Item)
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("item Saver: got item %d:%v", itemCount, item)
			itemCount++
			//parser.PrintProfile(item)
			//err := Save(item, client, index)

			//cal rpc to save item
			result := ""
			err = client.Call(config.ItemSaverRpc, item, &result)

			if err != nil {
				log.Printf("Item saver: error saving item %v: %v", item, err)
				continue
			}
		}
	}()
	return out, nil
}
