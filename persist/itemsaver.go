package persist

import (
	"context"
	"errors"
	"go-craler.com/engine"
	"log"
)
import "github.com/olivere/elastic/v7"

func ItemSaver(index string) (chan engine.Item, error) {
	out := make(chan engine.Item)
	client, err := elastic.NewClient(
		elastic.SetSniff(false), //部署在docker上面是内网，没办法sniff
	)
	if err != nil {
		return nil, err
	}
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("item Saver: got item %d:%v", itemCount, item)
			itemCount++
			//parser.PrintProfile(item)
			err := Save(item, client, index)
			if err != nil {
				log.Printf("Item saver: error saving item %v: %v", item, err)
				continue
			}
		}
	}()
	return out, nil
}

/**
保存数据到ES中
*/
func Save(item engine.Item, client *elastic.Client, index string) (err error) {

	if item.Type == "" {
		return errors.New("must supply type")
	}
	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, err = indexService.Do(context.Background()) // 兼容create update

	if err != nil {
		return err
	}

	return nil

}
