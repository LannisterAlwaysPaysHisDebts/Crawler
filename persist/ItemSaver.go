package persist

import (
	"context"
	"gopkg.in/olivere/elastic.v5"
)

const ES_INDEX = "dating_profile"
const ES_TYPE = "zhenai"

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		// todo: save
	}()
	return out
}

func Save(item interface{}) (id string, err error) {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return "", err
	}

	resp, err := client.Index().Index(ES_INDEX).
		Type(ES_TYPE).BodyJson(item).Do(context.Background())
	if err != nil {
		return "", err
	}

	return resp.Id, nil
}
