package engine

/**
  请求对象，
  1.需要访问的URL
  2. 解析的方法
*/
type Request struct {
	Url    string
	Parser Parser
}

type Parser interface {
	Parse(contents []byte, url string) ParseResult
	Serialize() (name string, args interface{})
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

type NilParser struct{}

func (n NilParser) Parse(contents []byte, url string) ParseResult {
	return ParseResult{}
}

func (n NilParser) Serialize() (name string, args interface{}) {
	return "NilParser", nil
}

type ParserFunc func(
	contents []byte, url string) ParseResult

type FuncParser struct {
	parser ParserFunc
	name   string
}

func (f *FuncParser) Parse(contents []byte, url string) ParseResult {
	return f.parser(contents, url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}

func NewFuncParser(parser ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: parser,
		name:   name,
	}
}
