package worker

import (
	"Crawler/engine"
	"Crawler/source/zhenAi/parser"
	"fmt"
	"log"
)

type SerializedParser struct {
	Name string
	Args interface{}
}

type Request struct {
	Url    string
	Parser SerializedParser
}

type ParserResult struct {
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

func SerializeResult(r engine.ParserResult) ParserResult {
	result := ParserResult{
		Items: r.Items,
	}

	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}
	return result
}

func DeserializeRequest(r Request) (engine.Request, error) {
	parser, err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url:    r.Url,
		Parser: parser,
	}, nil
}

func DeserializeResult(r ParserResult) (engine.ParserResult, error) {
	result := engine.ParserResult{
		Items: r.Items,
	}

	for _, req := range r.Requests {
		request, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("error request: %v", err)
			continue
		}
		result.Requests = append(result.Requests, request)
	}

	return result, nil
}

func deserializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case "ParserCityList":
		return engine.NewFuncParser(parser.ParserCityList, "ParserCityList"), nil
	case "ParserProfile":
		if userName, ok := p.Args.(string); ok {
			return parser.NewProfileParser(userName), nil
		} else {
			return nil, fmt.Errorf("error")
		}
	default:
		return nil, fmt.Errorf("error")
	}
}
