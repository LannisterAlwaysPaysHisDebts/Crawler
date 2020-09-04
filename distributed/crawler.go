package main

import (
	"Crawler/distributed/persist/client"
	"Crawler/engine"
	"Crawler/scheduler"
	"Crawler/source/zhenAi"
)

func main() {
	itemChan, err := client.ItemSaver(":1234")
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
