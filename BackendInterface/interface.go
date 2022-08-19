package BackendInterface

type Backend interface {
	SetListUrl string
	GetSetUrls([]string) []string
	GetCardUrlsFromSet(string, []string) string
	GetCardDataFromUrl(string) string
}