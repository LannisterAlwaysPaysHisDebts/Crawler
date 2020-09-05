package parser

import (
	"Crawler/engine"
	"regexp"
)

var (
	profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)" [^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

func ParserCity(contents []byte, _ string) engine.ParserResult {
	matches := profileRe.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ProfileParser(string(m[2])),
		})
	}

	matches2 := cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches2 {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParserCity,
		})
	}

	return result
}

func profileUrl(id string) string {
	return "https://album.zhenai.com/api/profile/getObjectProfile.do?objectID=" + id + "&data=eyJ2IjoiOWtubWpNYWVLZEJ6a2FiMnpYbzdRZz09Iiwib3MiOiJ3ZWIiLCJpdCI6MjkyLCJ0IjoiUktzdjRwcUw3V3hxRnpYMFhXdGVCREJBMDNNb3laWlhGQTJmVHd6SkRGWTM4dVlEYlhmMFJocTZIZnkyNmJLTit4dnREQVgxbWRlQ1FITGxEUHkwUXc9PSJ9&_=1599206047191&ua=h5%2F1.0.0%2F1%2F0%2F0%2F0%2F0%2F0%2F%2F0%2F0%2Fb032c96f-4048-417a-94d9-9af44ae697dc%2F0%2F0%2F294625501"
}
