package Web

import(
	"net/http"
	"golang.org/x/net/html"
	"time"
)


func FetchWebPage(url string) *html.Node {
	time.Sleep(1*time.Second)

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	rootNode, err := html.Parse(response.Body)
	if err != nil {
		panic(err)
	}

	return rootNode
}