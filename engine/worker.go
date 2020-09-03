package engine

import (
	"Crawler/fetcher"
	"log"
)

func Worker(r Request) (ParserResult, error) {
	log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetcher(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
		return ParserResult{}, err
	}

	return r.ParserFunc(body, r.Url), nil
}
