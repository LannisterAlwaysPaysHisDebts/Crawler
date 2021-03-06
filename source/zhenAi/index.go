package zhenAi

import (
	"Crawler/engine"
	"Crawler/source/zhenAi/parser"
)

const Index = "dating_profile"

func IndexRequest() engine.Request {
	return engine.Request{
		Url:    "http://www.zhenai.com/zhenghun/",
		Parser: engine.NewFuncParser(parser.ParserCityList, "ParserCityList"),
	}
}
