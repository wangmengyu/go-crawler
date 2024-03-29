package parser

import (
	"go-craler.com/engine"
	"go-craler.com/model"
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

//上海 1 | 32岁 2 | 大学本科 3 | 未婚 4 | 167cm 5 | 8001-12000元 6
const profileRe = `<div [^>]*class="des f-cl"[^>]*>([^<]+)</div>`
const numRe = `([0-9]+)`

var regex = regexp.MustCompile(profileRe)
var regexNum = regexp.MustCompile(numRe)
var regexId = regexp.MustCompile(`//album.zhenai.com/u/([\d]+)`)

func ParseProfile(bytes []byte, name string, gender string, url string) engine.ParseResult {
	matches := regex.FindAllSubmatch(bytes, -1)
	profile := model.Profile{Name: name, Gender: gender}
	results := engine.ParseResult{}
	var id string
	matchesId := regexId.FindAllSubmatch([]byte(url), -1)
	for _, match := range matchesId {
		id = string(match[1])
		break
	}

	for _, match := range matches {
		matchData := strings.Split(string(match[1]), "|")

		for i, m := range matchData {

			if i == 1 {
				profile.Age = getFirstIntFromStr([]byte(m))
			}

			if i == 4 {
				//height
				profile.Height = getFirstIntFromStr([]byte(m))
			}

			if i == 5 {
				profile.Income = m
			}

			//Xinzuo string 2
			if i == 3 {
				profile.Marriage = m
			}

			if i == 0 {
				profile.Hokou = m
			}
		}
	}

	results.Items = append(results.Items, engine.Item{
		Url:     url,
		Id:      id,
		Type:    "zhenai",
		Payload: profile,
	})

	return results

}

func PrintProfile(i interface{}) {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		log.Printf("[%v = %v]", f.Name, val)
	}
}

/**
从 string 抽取第一个整数数据
*/
func getFirstIntFromStr(bytes []byte) int {
	ageMatches := regexNum.FindAllSubmatch(bytes, -1)
	if len(ageMatches) > 0 {
		age := ageMatches[0][1]
		ageVal, err := strconv.Atoi(string(age))
		if err == nil {
			return ageVal
		}

	}
	return 0
}

type ProfileParser struct {
	name   string
	gender string
}

func (p *ProfileParser) Parse(contents []byte, url string) engine.ParseResult {
	return ParseProfile(contents, p.name, p.gender, url)
}

func (p *ProfileParser) Serialize() (name string, args interface{}) {
	profileMap := make(map[string]string)
	profileMap["name"] = p.name
	profileMap["gender"] = p.gender
	return "ParseProfile", profileMap
}

/**
  工厂方法 付初值
*/
func NewProfileParser(p map[string]interface{}) *ProfileParser {
	log.Printf("set profilr parser:%v", p)
	return &ProfileParser{name: p["name"].(string), gender: p["gender"].(string)}
}
