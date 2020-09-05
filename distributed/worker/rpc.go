package worker

import "Crawler/engine"

type CrawlService struct{}

func (CrawlService) Process(req Request, result *ParserResult) error {
	engineReq, err := DeserializeRequest(req)
	if err != nil {
		return err
	}

	engineResult, err := engine.Worker(engineReq)

	*result = SerializeResult(engineResult)
	return nil
}
