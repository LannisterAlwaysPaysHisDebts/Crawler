package parser

import (
	"Crawler/engine"
	"regexp"
)

var cityListRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+)" [^>]*>([^<]+)</a>`)

func ParserCityList(contents []byte) engine.ParserResult {
	matches := cityListRe.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}
	for _, m := range matches {
		// m[1]: url; m[2]: city
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParserCity,
		})
		break
	}

	return result
}
