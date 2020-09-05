package parser

import (
	"Crawler/engine"
	"regexp"
)

var cityListRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+)" [^>]*>([^<]+)</a>`)

func ParserCityList(contents []byte, _ string) engine.ParserResult {
	matches := cityListRe.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(m[1]),
			Parser: engine.NewFuncParser(ParserCity, "ParserCity"),
		})
		break
	}

	return result
}
