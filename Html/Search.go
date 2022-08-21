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

func FindNodesWithData(start *html.Node, targets []string, siblings bool) map[int]*html.Node {
	out := make(map[int]*html.Node)

	//is this our node
	for targetIndex, targetData := range targets {
		if start.Data == targetData {
			out[targetIndex] = start
			if len(out) == len(targets) {
				return out
			}
		}
	}
	//check our sibling (and it's children)
	if start.NextSibling != nil && siblings {
		found := FindNodesWithData(start.NextSibling, targets, true)
		for key, value := range found {
			out[key] = value
		}
		if len(out) == len(targets) {
			return out
		}
	}
	//check our children
	if start.FirstChild != nil {
		found := FindNodesWithData(start.FirstChild, targets, true)
		for key, value := range found {
			out[key] = value
		}
		if len(out) == len(targets) {
			return out
		}
	}

	return out
}
