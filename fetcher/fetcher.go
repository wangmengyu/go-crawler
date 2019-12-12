package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
)

/**
  获得指定url的HTML的UTF-8编码内容
*/
func Fetch(url string) ([]byte, error) {
	//获得城市列表页的HTML内容
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	//检查返回码
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Err: status code:", resp.StatusCode)
		return nil, fmt.Errorf("err: status code:%d", resp.StatusCode)
	}

	//获得当前页面的编码
	bodyReader := bufio.NewReader(resp.Body)
	e, err := determineEncoding(bodyReader)
	if err != nil {
		return nil, err
	}
	//将当前页面编码转换为utf8
	utf8Body := transform.NewReader(bodyReader, e.NewDecoder())

	//读取UTF8编码的HTML内容
	return ioutil.ReadAll(utf8Body)

}

func determineEncoding(r *bufio.Reader) (encoding.Encoding, error) {
	bytes, err := r.Peek(1024)
	if err != nil {
		return nil, err
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e, nil
}
