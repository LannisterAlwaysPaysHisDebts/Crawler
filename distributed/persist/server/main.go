package main

import (
	"Crawler/config"
	"Crawler/distributed/persist"
	"Crawler/distributed/rpcsupport"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func main() {
	log.Fatal(serveRpc(config.RPC_PORT, "dating_profile"))
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}
	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
}
