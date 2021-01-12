package client

import (
	"Crawler/config"
	"Crawler/distributed/worker"
	"Crawler/engine"
	"net/rpc"
)

// 创建request处理器
func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {
	// rpc调用CrawlServiceRpc方法解析request
	return func(r engine.Request) (result engine.ParserResult, e error) {
		sReq := worker.SerializeRequest(r)
		var sResult worker.ParserResult

		c := <-clientChan
		err := c.Call(config.CrawlServiceRpc,
			sReq, &sResult)
		if err != nil {
			return engine.ParserResult{}, err
		}

		return worker.DeserializeResult(sResult)
	}
}
