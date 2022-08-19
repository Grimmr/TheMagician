package Web

import(
	"net/http"
	"golang.org/x/net/html"
	"time"
)

var limit int64 = 1000000000
var lastFetch int64 = 0 

func FetchWebPage(url string) *html.Node {
	now := time.Now().UnixNano()

	if (now - lastFetch) < limit {
		time.Sleep(time.Duration(now-lastFetch))
	}

	lastFetch = now

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