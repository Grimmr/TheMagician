package YGO

import (
	"github.com/Grimmr/TheMagician/Html"
	"golang.org/x/net/html"
)

func (this *YgoBackend) GetCardUrlsFromSet(url string) []string {
	//tableUrl := getTableUrlFromHtml(Web.FetchWebPage(url))
	return []string{}
}

func (this *YgoBackend) getTableUrlFromHtml(node *html.Node) string {
	naviNode := Html.FindNodesWithAttrs(node, []map[string]string{{"class": "set-navigation"}})[0]

	foundLinkNodes := Html.FindNodesWithData(naviNode, []string{"European English"})
	if len(foundLinkNodes) == 0 { //we didn't find the EU english
		foundLinkNodes = Html.FindNodesWithData(naviNode, []string{"North American English"}) //default to NA
	}
	linkNode := foundLinkNodes[0]

	urlSuffix := linkNode.Parent.Attr[0].Val

	return this.GetUrlPrefix() + urlSuffix
}
