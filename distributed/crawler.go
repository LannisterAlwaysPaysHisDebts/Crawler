package main

import (
	"Crawler/config"
	"Crawler/distributed/persist/client"
	"Crawler/engine"
	"Crawler/scheduler"
	"Crawler/source/zhenAi"
)

func main() {
	itemChan, err := client.ItemSaver(config.RPC_PORT)
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan:    itemChan,
	}

	e.Run(zhenAi.IndexRequest())
}
