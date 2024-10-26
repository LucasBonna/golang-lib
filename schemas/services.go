package schemas

type ServiceNames string

const (
	CFCrawler       ServiceNames = "cfcrawler"
	CFStorage       ServiceNames = "cfstorage"
	CFNotifications ServiceNames = "cfnotifications"
)

var ServiceUrls = map[ServiceNames]string{
	CFCrawler:       "http://cfcrawler-api",
	CFStorage:       "http://localhost:8001",
	CFNotifications: "http://cfcnotifications-api",
}
