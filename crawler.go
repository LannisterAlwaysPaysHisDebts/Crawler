package main

import (
	"Crawler/engine"
	"Crawler/persist"
	"Crawler/scheduler"
	"Crawler/source/zhenAi"
)

func main() {
	itemChan, err := persist.ItemSaver()
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
