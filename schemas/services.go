package schemas

type ServiceNames string

const (
	CFGateway       ServiceNames = "cfgateway"
	CFCrawler       ServiceNames = "cfcrawler"
	CFStorage       ServiceNames = "cfstorage"
	CFNotifications ServiceNames = "cfnotifications"
)

var ServiceUrls = map[ServiceNames]string{
	CFGateway:       "http://cfgateway-api:5500",
	CFCrawler:       "http://cfcrawler-api:8002",
	CFStorage:       "http://cfstorage-api:8001",
	CFNotifications: "http://cfcnotifications-api:8000",
}
