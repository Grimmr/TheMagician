package YGO

import (
	"github.com/Grimmr/TheMagician/Html"
	"github.com/Grimmr/TheMagician/Web"
	"golang.org/x/net/html"
)

func (this YgoBackend) GetCardUrlsFromSet(url string) []string {
	tableUrl := this.getTableUrlFromHtml(Web.FetchWebPage(url))
	return this.getCardUrlsFromTableHtml(Web.FetchWebPage(tableUrl))
}

func (this *YgoBackend) getTableUrlFromHtml(node *html.Node) string {
	naviRowNode := Html.FindNodesWithAttrs(node, []map[string]string{{"class": "set-navigation"}})[0].FirstChild.NextSibling

	foundLinkNodes := Html.FindNodesWithData(naviRowNode, []string{"European English", "North American English", "English"}, false)
	linkNode := foundLinkNodes[0]
	if linkNode == nil { //we didn't find the EU english
		linkNode = foundLinkNodes[1]
	}
	if linkNode == nil { //we didn't find the NA english
		linkNode = foundLinkNodes[2]
	}

	urlSuffix := linkNode.Parent.Attr[0].Val

	return this.GetUrlPrefix() + urlSuffix
}

func (this *YgoBackend) getCardUrlsFromTableHtml(node *html.Node) []string {
	out := make([]string, 0)

	setListTableNode := Html.FindNodesWithAttrs(node, []map[string]string{{"class": "set-list"}})[0]
	bodyNode := setListTableNode.FirstChild.NextSibling.FirstChild
	currentRow := bodyNode.FirstChild.NextSibling
	for currentRow != nil {
		targetCol := currentRow.FirstChild.NextSibling
		targetChild := Html.FindNodesWithData(targetCol, []string{"a"}, false)[0]
		out = append(out, this.GetUrlPrefix()+targetChild.Attr[0].Val)
		currentRow = currentRow.NextSibling
	}

	return out
}
