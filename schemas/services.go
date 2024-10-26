package schemas

type ServiceNames string

const (
	CFCrawler       ServiceNames = "cfcrawler"
	CFStorage       ServiceNames = "cfstorage"
	CFNotifications ServiceNames = "cfnotifications"
)

var ServiceUrls = map[ServiceNames]string{
	CFCrawler:       "http://cfcrawler-api",
	CFStorage:       "http://cfstorage-api",
	CFNotifications: "http://cfcnotifications-api",
}
