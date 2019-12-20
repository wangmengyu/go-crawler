package config

const (
	// service port
	ItemSaverPort = 1234
	WorkderPort0  = 9000

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
