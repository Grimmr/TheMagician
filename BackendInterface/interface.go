package BackendInterface

type Backend interface {
	GetUrlPrefix() string
	GetSetListUrl() string
	GetSetUrls(string) []map[string][]string     //url onto list of map of url onto list of names
	GetCardUrlsFromSet(string) []string          //url onto list of url
	GetCardDataFromUrl(string) map[string]string //url onto card data (as map of field onto data)
}
