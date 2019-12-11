package parser

import (
	"go-craler.com/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(bytes []byte) engine.ParseResult {
	results := engine.ParseResult{}
	regex := regexp.MustCompile(cityListRe)
	matches := regex.FindAllSubmatch(bytes, -1)

	for _, match := range matches {

		url := match[1]
		city := match[2]
		results.Requests = append(
			results.Requests,
			engine.Request{
				Url:        string(url),
				ParserFunc: engine.NilParse,
			})
		results.Items = append(results.Items, string(city))
	}

	return results

}
