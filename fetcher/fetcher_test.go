package fetcher

import (
	"fmt"
	"testing"
)

func TestFetcher(t *testing.T) {
	url := "https://album.zhenai.com/u/1111994553"

	result, err := Fetcher(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", result)
}
