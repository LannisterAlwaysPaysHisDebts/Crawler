package main

import (
	"Crawler/distributed/persist/client"
	"Crawler/distributed/rpcsupport"
	worker "Crawler/distributed/worker/client"
	"Crawler/engine"
	"Crawler/scheduler"
	"Crawler/source/zhenAi"
	"flag"
	"log"
	"net/rpc"
	"strings"
)

var (
	itemSaverHost = flag.String("itemsaver_host", "", "itemsaver host")
	workerHosts   = flag.String("worker_hosts", "", "worker hosts  (comma separated)")
)

func main() {
	itemChan, err := client.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}

	processor := worker.CreateProcessor(
		createClientPool(
			strings.Split(*workerHosts, ",")))

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      10,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}

	e.Run(zhenAi.IndexRequest())
}

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client

	for _, h := range hosts {
		c, err := rpcsupport.NewClient(h)
		if err != nil {
			log.Printf("")
		} else {
			clients = append(clients, c)
			log.Printf("")
		}
	}

	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, c := range clients {
				out <- c
			}
		}
	}()
	return out
}
