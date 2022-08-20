package YGO

import (
	"github.com/Grimmr/TheMagician/Web"
	"golang.org/x/net/html"
)

type YgoBackend struct {
}

func (this *YgoBackend) GetCardDataFromUrl(url string) map[string]string {
	return getCardDataFromHtml(Web.FetchWebPage(url))
}

func getCardDataFromHtml(root *html.Node) map[string]string {
	return make(map[string]string)
}
