package main

import (
	"Crawler/config"
	"Crawler/distributed/rpcsupport"
	"Crawler/engine"
	"Crawler/model"
	"fmt"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	// start server
	go serveRpc(fmt.Sprintf(":%d", config.RpcPort), "test1")
	time.Sleep(time.Second)

	// start client
	client, err := rpcsupport.NewClient(fmt.Sprintf(":%d", config.RpcPort))
	if err != nil {
		panic(err)
	}

	// call save
	item := engine.Item{
		Url:  "https://album.zhenai.com/u/1204455387",
		Type: "zhenai",
		Id:   "1204455387",
		Payload: model.Profile{
			Name:       "夏洛克",
			Gender:     "女",
			Age:        23,
			Height:     163,
			Weight:     57,
			Income:     "5001-8000元",
			Marriage:   "未婚",
			Education:  "大学本科",
			Occupation: "建筑师",
			Hokou:      "建筑师",
			Xinzou:     "天秤座",
			House:      "租房",
			Car:        "未买车",
		},
	}

	var result string
	err = client.Call(config.ItemSaverRpc, item, &result)
	if err != nil || result != "ok" {
		t.Errorf("result %s; error : %s", result, err)
	}
}
