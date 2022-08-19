package BackendInterface

type Backend interface {
	GetSetListUrl() string
	GetSetUrls([]string) []string
	GetCardUrlsFromSet(string, []string) string
	GetCardDataFromUrl(string) map[string]string
}