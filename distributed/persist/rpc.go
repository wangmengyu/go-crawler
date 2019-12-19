package persist

import (
	"github.com/olivere/elastic/v7"
	"go-craler.com/engine"
	"go-craler.com/persist"
	"log"
)

/**
  持久化存储服务结构
*/
type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	err := persist.Save(item, s.Client, s.Index)
	if err == nil {
		*result = "ok"
		log.Printf("item %v saved", item)
	} else {
		log.Printf("err saving item %v, error %v", item, err)
	}
	return err
}
