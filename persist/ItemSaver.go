package persist

import (
	"Crawler/engine"
	"context"
	"errors"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

/*
获取一个Item channel, 同时注册一个goroutine用来保存channel传递过来的数据
@index es保存的index名称
*/
func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			itemCount++
			err := Save(index, client, item)
			if err != nil {
				log.Printf("Item Saver: error saving item %v: %v", item, err)
			}
		}
	}()
	return out, nil
}

// 保存数据
func Save(index string, client *elastic.Client, item engine.Item) error {
	if item.Type == "" {
		return errors.New("must supply Type")
	}

	indexService := client.Index().Index(index).
		Type(item.Type).
		BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, err := indexService.Do(context.Background())
	if err != nil {
		return err
	}

	return nil
}
