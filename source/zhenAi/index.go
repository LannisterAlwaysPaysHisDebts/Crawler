package zhenAi

import (
	"Crawler/engine"
	"Crawler/source/zhenAi/parser"
)

func IndexRequest() engine.Request {
	return engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/",
		ParserFunc: parser.ParserCityList,
	}
}
