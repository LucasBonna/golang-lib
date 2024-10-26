package schemas

type ServiceNames string

const (
	CFGateway       ServiceNames = "cfgateway"
	CFCrawler       ServiceNames = "cfcrawler"
	CFStorage       ServiceNames = "cfstorage"
	CFNotifications ServiceNames = "cfnotifications"
)

var ServiceUrls = map[ServiceNames]string{
	CFGateway:       "http://cfgateway-api",
	CFCrawler:       "http://cfcrawler-api",
	CFStorage:       "http://cfstorage-api",
	CFNotifications: "http://cfcnotifications-api",
}
