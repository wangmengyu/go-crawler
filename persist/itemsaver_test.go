package persist

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"go-craler.com/engine"
	"go-craler.com/model"
	"testing"
)

func TestSave(t *testing.T) {
	expeted := engine.Item{
		Url:  "https://album.zhenai.com/u/1486293757",
		Id:   "1486293757",
		Type: "zhenai",
		Payload: model.Profile{
			Name:       "维E真子",
			Gender:     "女士",
			Age:        54,
			Height:     158,
			Weight:     0,
			Income:     "20001-50000元",
			Marriage:   "离异",
			Occupation: "",
			Hokou:      "上海",
			Xinzuo:     "",
			House:      "",
			Car:        "",
		},
	}
	client, err := elastic.NewClient(
		elastic.SetSniff(false), //部署在docker上面是内网，没办法sniff
	)
	index := "dating_test"
	err = Save(expeted, client, index)
	if err != nil {
		panic(err)
	}
	resp, err := client.Get().
		Index(index).
		Type(expeted.Type).Id(expeted.Id).Do(context.Background())
	if err != nil {
		panic(err)
	}

	var actual engine.Item
	err = json.Unmarshal(resp.Source, &actual)

	if err != nil {
		panic(err)
	}

	actualProfile, err := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	if actual != expeted {
		t.Errorf("got %v; expeted %v", actual, expeted)

	}

}
