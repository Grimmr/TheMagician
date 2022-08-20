package YGO

import (
	"github.com/Grimmr/TheMagician/Html"
	"github.com/Grimmr/TheMagician/Web"
	"golang.org/x/net/html"
)

type YgoBackend struct {
}

func (this *YgoBackend) GetCardDataFromUrl(url string) map[string]string {
	return getCardDataFromHtml(Web.FetchWebPage(url))
}

func getCardDataFromHtml(root *html.Node) map[string]string {
	out := make(map[string]string)

	targets := []map[string]string{
		{"title": "Password"},  //password title node
		{"id": "firstHeading"}, //name
		{"title": "Card type"}} //card type title node

	found := Html.FindNodesWithAttrs(root, targets)

	pwNode := Html.FindNodesWithAttrs(found[0].Parent.Parent, []map[string]string{{"class": "mw-redirect"}})[0]
	out["password"] = pwNode.FirstChild.Data

	out["name"] = found[1].FirstChild.Data

	typeNode := Html.FindNodesWithAttrs(found[2].Parent.Parent, []map[string]string{{"class": ""}})[0]
	out["type"] = typeNode.FirstChild.NextSibling.FirstChild.Data

	return out
}
