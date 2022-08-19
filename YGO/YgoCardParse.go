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

func findInfoColumn(start *html.Node) *html.Node {
	//is this our node
	for _, attr := range start.Attr {
		if attr.Key == "class" && attr.Val == "infocolumn" {
			return start
		}
	}
	//check our sibling (and it's children)
	if start.NextSibling != nil {
		found := findInfoColumn(start.NextSibling)
		if found != nil {
			return found
		}
	}
	//check our children
	if start.FirstChild != nil {
		found := findInfoColumn(start.FirstChild)
		if found != nil {
			return found
		}
	}

	//couldn't find it
	return nil
}