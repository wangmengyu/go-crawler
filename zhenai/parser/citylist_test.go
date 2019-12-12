package parser

import (
	"io/ioutil"
	"testing"
)

/**
  测试城市列表解析器
*/
func TestParseCityList(t *testing.T) {
	//先用fetch方法获取HTML内容，复制出来保存到citylist_test_data.html
	/*
		contests,err := fetcher.Fetch("http://www.zhenai.com/zhenghun/")
		if err!=nil {
			panic(err)
		}
	*/
	//从文件里读取contents
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}
	results := ParseCityList(contents)

	const resultSize = 470 //记录总数验证
	if len(results.Requests) != resultSize {
		t.Errorf("result should have %d requests but had %d", resultSize, len(results.Requests))
	}

	//希望的3个url & city
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	expectedCities := []string{"City 阿坝", "City 阿克苏", "City 阿拉善盟"}

	//
	for i, url := range expectedUrls {
		if url != results.Requests[i].Url {
			t.Errorf("expeted URL #%d %s , but was %s", i, url, results.Requests[i].Url)
		}
	}

	if len(results.Items) != resultSize {
		t.Errorf("result should have %d Items but had %d", resultSize, len(results.Items))
	}

	for i, city := range expectedCities {
		if city != results.Items[i].(string) {
			t.Errorf("expeted city #%d %s , but was %s", i, city, results.Items[i].(string))
		}
	}
	//fmt.Printf("%s\n", contents)

}
