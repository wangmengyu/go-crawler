package engine

/**
  请求对象，
  1.需要访问的URL
  2. 解析的方法
*/
type Request struct {
	Url        string
	ParserFunc func(contents []byte) ParseResult
}

/**
  解析器返回的结果：
   1. 下一级的更多请求，
   2. 解析出来的结果明细
*/
type ParseResult struct {
	Requests []Request
	Items    []Item // interface{} 代表任何类型

}

type Item struct {
	Url     string
	Id      string
	Type    string
	Payload interface{}
}

func NilParse([]byte) ParseResult {
	return ParseResult{}
}
