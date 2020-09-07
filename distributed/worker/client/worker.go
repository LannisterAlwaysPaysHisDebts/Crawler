package client

import (
	"Crawler/config"
	"Crawler/distributed/rpcsupport"
	"Crawler/distributed/worker"
	"Crawler/engine"
	"fmt"
)

func CreateProcessor() (engine.Processor, error) {
	client, err := rpcsupport.NewClient(fmt.Sprintf(":%d", config.RpcPort))
	if err != nil {
		return nil, err
	}

	return func(r engine.Request) (result engine.ParserResult, e error) {
		sReq := worker.SerializeRequest(r)
		var sResult worker.ParserResult

		err = client.Call(config.CrawlServiceRpc,
			sReq, &sResult)
		if err != nil {
			return engine.ParserResult{}, err
		}

		return worker.DeserializeResult(sResult)
	}, nil
}
