package parser

import (
	"fmt"
	"go-craler.com/engine"
	"go-craler.com/model"
	"log"
	"regexp"
	"strconv"
)

const purpleRe = `class="m-btn purple"[^>]*>([^<]+)<`
const pinkRe = `class="m-btn pink"[^>]*>([^<]+)<`

var regex = regexp.MustCompile(purpleRe)
var regexPink = regexp.MustCompile(pinkRe)
var regexNum = regexp.MustCompile(`([0-9]+)`)

func ParseProfile(bytes []byte) engine.ParseResult {

	matches := regex.FindAllSubmatch(bytes, -1)
	matchesPink := regexPink.FindAllSubmatch(bytes, -1)
	profile := model.Profile{}
	/**
	Age int  2
	Height int  4
	Weight int 5
	Income string 6
	Marriage string 1
	Occupation string 8
	Hokou string 11
	Xinzuo string 3
	House string 15
	Car string 16
	*/
	results := engine.ParseResult{}

	for i, match := range matches {

		/**
		Age int  1
		Height int  3
		Weight int 4
		Income string 6
		Marriage string 1
		Occupation string 7
		Xinzuo string 2
		*/
		if i == 0 {
			profile.Marriage = string(match[1])
			log.Printf("婚配:%s", profile.Marriage)
		}
		if i == 1 {
			//age
			profile.Age = getFirstIntFromStr(match[1])
			log.Printf("年纪:%d", profile.Age)
		}
		if i == 3 {
			//height
			profile.Height = getFirstIntFromStr(match[1])
			log.Printf("H:%d", profile.Height)
		}
		if i == 4 {
			//height
			profile.Weight = getFirstIntFromStr(match[1])
			log.Printf("W:%d", profile.Weight)
		}
		if i == 6 {
			profile.Income = string(match[1])
			log.Printf("收入:%d", profile.Weight)
		}

		//Occupation string 7
		if i == 7 {
			profile.Occupation = string(match[1])
			log.Printf("职位:%s", profile.Occupation)
		}
		//Xinzuo string 2
		if i == 2 {
			profile.Xinzuo = string(match[1])
			log.Printf("星座:%s", profile.Xinzuo)
		}

	}

	for i, match := range matchesPink {
		if i == 1 {
			//hukou
			profile.Hokou = string(match[1])
			log.Printf("户口:%s", profile.Hokou)
		}
		if i == 5 {
			//Car
			profile.House = string(match[1])
			log.Printf("房子:%s", profile.House)
		}

		if i == 6 {
			//hukou
			profile.Car = string(match[1])
			log.Printf("房子:%s", profile.Car)
		}
	}

	fmt.Println("profile=", profile)
	results.Items = append(results.Items, profile)

	return results

}

/**
从 string 抽取第一个整数数据
*/
func getFirstIntFromStr(bytes []byte) int {
	ageMatches := regexNum.FindAllSubmatch(bytes, -1)
	age := ageMatches[0][1]
	ageVal, err := strconv.Atoi(string(age))
	if err == nil {
		return ageVal
	}
	return 0
}
