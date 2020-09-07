package main

import (
	"Crawler/config"
	"Crawler/distributed/persist/client"
	worker "Crawler/distributed/worker/client"
	"Crawler/engine"
	"Crawler/scheduler"
	"Crawler/source/zhenAi"
	"fmt"
)

func main() {
	itemChan, err := client.ItemSaver(fmt.Sprintf(":%d", config.RpcPort))
	if err != nil {
		panic(err)
	}

	processor, err := worker.CreateProcessor()
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      10,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}

	e.Run(zhenAi.IndexRequest())
}
