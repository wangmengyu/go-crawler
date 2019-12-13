package parser

import (
	"fmt"
	"go-craler.com/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(bytes []byte) engine.ParseResult {
	results := engine.ParseResult{}
	regex := regexp.MustCompile(cityListRe)
	matches := regex.FindAllSubmatch(bytes, -1)

	limit := 10
	for _, match := range matches {
		fmt.Println("match:", match)
		url := match[1]
		city := match[2]
		fmt.Printf("url:%s, city:%s\n", string(url), string(city))
		results.Requests = append(
			results.Requests,
			engine.Request{
				Url:        string(url),
				ParserFunc: ParseCity,
			})
		results.Items = append(results.Items, "City "+string(city))
		limit--
		if limit == 0 {
			break
		}

	}

	return results

}
