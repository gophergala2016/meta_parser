package main

import (
	"golang.org/x/net/html"
)

func MetaParser(doc *html.Node) map[string]string {
	metas := map[string]string{}
	var ParseMetaAttr func(*html.Node)
	ParseMetaAttr = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "meta" {
			name, content := "", ""
			for _, a := range n.Attr {
				if a.Key == "name" {
					name = a.Val
				}
				if a.Key == "content" {
					content = a.Val
				}
			}
			if name != "" && content != "" {
				metas[name] = content
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			ParseMetaAttr(c)
		}
	}
	ParseMetaAttr(doc)
	return metas
}
