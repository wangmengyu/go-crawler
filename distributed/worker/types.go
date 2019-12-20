package worker

import (
	"errors"
	"fmt"
	"go-craler.com/distributed/config"
	"go-craler.com/engine"
	"go-craler.com/zhenai/parser"
	"log"
	"reflect"
)

/**
  序列化的解析器
*/
type SerializedParser struct {
	Name string      //方法名
	Args interface{} // 参数
}

type Request struct {
	Url    string
	Parser SerializedParser
}
type ParseResult struct {
	Items    []engine.Item
	Requests []Request
}

func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

func SerializeResult(r engine.ParseResult) ParseResult {
	result := ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}
	return result
}

func DeserializeRequest(r Request) (engine.Request, error) {
	p, err := DeserializeParser(r.Parser)
	if err == nil {
		return engine.Request{
			Url:    r.Url,
			Parser: p,
		}, nil
	} else {
		return engine.Request{}, err
	}
}

func DeserializeResult(r ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items: r.Items,
	}
	for _, r := range r.Requests {
		req, err := DeserializeRequest(r)
		if err != nil {
			log.Printf("err deserializing request %v", err)
			continue
		} else {
			result.Requests = append(result.Requests, req)
		}
	}
	return result

}

func DeserializeParser(p SerializedParser) (engine.Parser, error) {
	//return engine.Parser()
	switch p.Name {
	case config.ParseCityList:
		return engine.NewFuncParser(parser.ParseCityList, config.ParseCityList), nil
	case config.ParseCity:
		return engine.NewFuncParser(parser.ParseCity, config.ParseCity), nil
	case config.ParseProfile:
		fmt.Println("p-name:", p.Name)
		fmt.Println("p-args:", reflect.TypeOf(p.Args))
		if profile, ok := p.Args.(map[string]interface{}); ok {
			return parser.NewProfileParser(profile), nil
		} else {
			return nil, fmt.Errorf("invalid arg:%v", p.Args)
		}
	case config.NilParser:
		return engine.NilParser{}, nil
	default:
		return nil, errors.New("unknown parser name")

	}

}
