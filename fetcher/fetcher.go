package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

/**
  获得指定url的HTML的UTF-8编码内容
*/
var rateLimiter = time.Tick(1000 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	//获得城市列表页的HTML内容
	<-rateLimiter //限速
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
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
