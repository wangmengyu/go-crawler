package parser

import (
	"fmt"
	"go-craler.com/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(bytes []byte) engine.ParseResult {
	results := engine.ParseResult{}
	regex := regexp.MustCompile(cityRe)
	matches := regex.FindAllSubmatch(bytes, -1)

	for _, match := range matches {
		fmt.Println("match:", match)
		url := match[1]
		user := match[2]
		fmt.Printf("url:%s, user:%s\n", string(url), string(user))
		results.Requests = append(
			results.Requests,
			engine.Request{
				Url:        string(url),
				ParserFunc: ParseProfile,
			})
		results.Items = append(results.Items, "User "+string(user))

	}

	return results

}
