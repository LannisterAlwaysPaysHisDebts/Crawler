package zhenAi

import (
	"myGit/Crawler/engine"
	"myGit/Crawler/source/zhenAi/parser"
)

func IndexRequest() engine.Request {
	return engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/",
		ParserFunc: parser.ParserCityList,
	}
}
