package engine

import (
	"log"
	"myGit/Crawler/fetcher"
)

type SimpleEngine struct{}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parserResult, err := Worker(r)
		if err != nil {
			continue
		}

		requests = append(requests, parserResult.Requests...)

		for _, item := range parserResult.Items {
			log.Printf("Got item: %+v", item)
		}
	}
}

func Worker(r Request) (ParserResult, error) {
	log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetcher(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
		return ParserResult{}, err
	}

	return r.ParserFunc(body), nil
}