package main

import (
	"Crawler/distributed/rpcsupport"
	"Crawler/distributed/worker"
	"fmt"
	"log"
)

const WorkerPort0 = 9999

func main() {
	// work port
	log.Fatal(rpcsupport.ServeRpc(
		fmt.Sprintf(":%d", WorkerPort0),
		worker.CrawlService{}))
}
