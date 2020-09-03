package parser

import (
	"Crawler/engine"
	"Crawler/model"
	"regexp"
	"strconv"
)

var CommonCompile = regexp.MustCompile(`<div class="des f-cl" data-v-3c42fade>([^\|]+) \| ([\d]+)岁 \| ([^\|]+) \| ([^\|]+) \| ([\d]+)cm \| ([\d]+-[\d]+)元</div>`)

var idUrlRe = regexp.MustCompile(`https://album.zhenai.com/u/([\d]+)`)

// todo: test, 猜你喜欢
func ParserProfile(contents []byte, url string, name string) engine.ParserResult {
	profile := model.Profile{Name: name}

	match := CommonCompile.FindSubmatch(contents)
	if match != nil {
		profile.Hokou = string(match[1])
		profile.Education = string(match[3])
		profile.Marriage = string(match[4])
		profile.Income = string(match[6])

		age, err := strconv.Atoi(string(match[2]))
		if err == nil {
			profile.Age = age
		}

		height, err := strconv.Atoi(string(match[5]))
		if err == nil {
			profile.Height = height
		}
	}

	return engine.ParserResult{
		Items: []engine.Item{
			{
				Url:     url,
				Type:    "zhenai",
				Id:      extractString([]byte(url), idUrlRe),
				Payload: profile,
			},
		},
	}
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
