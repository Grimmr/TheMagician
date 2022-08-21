package YGO

import (
	"time"

	"github.com/Grimmr/TheMagician/Html"
	"github.com/Grimmr/TheMagician/Web"
	"golang.org/x/net/html"
)

func (this YgoBackend) GetSetUrls(url string) []map[string]string {
	return this.getSetUrlsFromHtml(Web.FetchWebPage(url))
}

func (this *YgoBackend) getSetUrlsFromHtml(node *html.Node) []map[string]string {
	out := make([]map[string]string, 0)

	tcgTextNode := Html.FindNodesWithAttrs(node, []map[string]string{{"id": "TCG"}})[0]

	currentNode := tcgTextNode.Parent
	year := ""
	month := ""
	day := ""
listloop:
	for currentNode != nil {
		//current node is year
		if currentNode.Data == "h3" {
			yearNode := Html.FindNodesWithAttrs(currentNode, []map[string]string{{"class": "mw-headline"}})[0]
			year = yearNode.FirstChild.Data
		}

		//current node is month
		if currentNode.Data == "h4" {
			monthNode := Html.FindNodesWithAttrs(currentNode, []map[string]string{{"class": "mw-headline"}})[0]
			month = monthNode.FirstChild.Data
		}

		//current node is list of sets
		if currentNode.Data == "ul" {
			currentSubNode := currentNode.FirstChild
			for currentSubNode != nil {
				if currentSubNode.Data == "li" {
					day = currentSubNode.FirstChild.Data[:len(currentSubNode.FirstChild.Data)-3]
					if len(day) == 1 {
						day = "0" + day
					}

					//stop if we are in the future
					date := year + "-" + month + "-" + day
					dateAsObject, _ := time.Parse("2006-January-02", date)
					if time.Now().Before(dateAsObject) {
						break listloop
					}

					//create rows
					currentSubSubNode := currentSubNode.FirstChild.NextSibling
					for currentSubSubNode != nil {
						if currentSubSubNode.Data == "a" {
							name := currentSubSubNode.FirstChild.Data
							urlSuffix := currentSubSubNode.Attr[0].Val
							row := map[string]string{"name": name, "date": date, "url": this.GetUrlPrefix() + urlSuffix}
							out = append(out, row)

						}
						currentSubSubNode = currentSubSubNode.NextSibling
					}
				}
				currentSubNode = currentSubNode.NextSibling
			}
		}

		currentNode = currentNode.NextSibling
	}
	return out
}
