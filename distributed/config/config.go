package config

const (
	// rate limit
	Qps = 1

	//es table name
	ElasticIndex = "dating_profile"
	//rpc endpoints
	ItemSaverRpc    = "ItemSaverService.Save"
	CrawlServiceRpc = "CrawlService.Process"

	ParseCity     = "ParseCity"
	ParseCityList = "ParseCityList"
	ParseProfile  = "ParseProfile"
	NilParser     = "NilParser"
)
