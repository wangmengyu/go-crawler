package worker

import "go-craler.com/engine"

type CrawlService struct {
}

func (CrawlService) Process(req Request, result *ParseResult) error {
	engineReq, err := DeserializeRequest(req)
	if err != nil {
		return err
	}

	eResult, err := engine.Worker(engineReq)
	if err != nil {
		return err
	}

	*result = SerializeResult(eResult)
	return nil
}
