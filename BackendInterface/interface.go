package BackendInterface

type Backend interface {
	GetUrlPrefix() string
	GetSetListUrl() string
	GetSetUrls(string) []string
	GetCardUrlsFromSet(string) []string
	GetCardDataFromUrl(string) map[string]string
}
