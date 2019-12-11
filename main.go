package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	//获得城市列表页的HTML内容
	resp, err := http.Get("http://www.zhenai.com/zhenghun/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//检查返回码
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Err: status code:", resp.StatusCode)
		return
	}

	//获得当前页面的编码
	e := determineEncoding(resp.Body)
	//将当前页面编码转换为utf8
	utf8Body := transform.NewReader(resp.Body, e.NewDecoder())

	//读取UTF8编码的HTML内容
	all, err := ioutil.ReadAll(utf8Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", all)

}

/**
  get encoding of an io.Reader
*/
func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
