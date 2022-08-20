package Html

import (
	"golang.org/x/net/html"
)

func FindNodesWithAttrs(start *html.Node, targets []map[string]string) map[int]*html.Node {
	out := make(map[int]*html.Node)

	//is this our node
	for targetIndex, targetAttrSet := range targets {
		hasAll := true
		for key, value := range targetAttrSet {
			found := false
			for _, checkAttr := range start.Attr {
				if key == checkAttr.Key && value == checkAttr.Val {
					found = true
					break
				}
			}
			if !found {
				hasAll = false
				break
			}
		}
		if hasAll {
			out[targetIndex] = start
		}
	}
	//check our sibling (and it's children)
	if start.NextSibling != nil {
		found := FindNodesWithAttrs(start.NextSibling, targets)
		for key, value := range found {
			out[key] = value
		}
	}
	//check our children
	if start.FirstChild != nil {
		found := FindNodesWithAttrs(start.FirstChild, targets)
		for key, value := range found {
			out[key] = value
		}
	}

	return out
}
