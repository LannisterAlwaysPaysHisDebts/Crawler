package fetcher

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Fetcher(url string) ([]byte, error) {
	log.Printf("Fetching url %s", url)

	client := http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.125 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)
}
