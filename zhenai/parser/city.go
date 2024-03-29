package parser

import (
	"fmt"
	"go-craler.com/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

var regexCity = regexp.MustCompile(cityRe)

const genderRe = `<span class="grayL">性别：</span>([^<]+)</td>`

var regexGen = regexp.MustCompile(genderRe)
var cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/shanghai/[^"]+)">`)

func ParseCity(bytes []byte, _ string) engine.ParseResult {
	results := engine.ParseResult{}
	matches := regexCity.FindAllSubmatch(bytes, -1)
	matchesGen := regexGen.FindAllSubmatch(bytes, -1)
	for i, match := range matches {
		//fmt.Println("match:", match)
		url := match[1]
		user := match[2] // 此处必须深度拷贝。因为该参数回传递给后续的抓取，只有深度拷贝才不会调用match.
		gender := matchesGen[i][1]
		fmt.Printf("url:%s, user:%s\n", string(url), string(user))
		profileMap := make(map[string]interface{})
		profileMap["name"] = string(user)
		profileMap["gender"] = string(gender)
		results.Requests = append(
			results.Requests,
			engine.Request{
				Url:    string(url),
				Parser: NewProfileParser(profileMap)})
		//results.Items = append(results.Items, "User "+string(user))

	}
	matches = cityUrlRe.FindAllSubmatch(bytes, -1)
	for _, m := range matches {
		results.Requests = append(results.Requests, engine.Request{
			Url:    string(m[1]),
			Parser: engine.NewFuncParser(ParseCity, "ParseCity"),
		})
	}

	return results

}
