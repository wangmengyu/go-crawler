package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile.html")
	if err != nil {
		return
	}

	//fmt.Println(string(contents))

	results := ParseProfile(contents, "test", "test", "")

	fmt.Printf("%v", results)

}
