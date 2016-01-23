package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://github.com/")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	metas := map[string]string{}
	var f func(*html.Node)
	f = func(n *html.Node) {
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
			f(c)
		}
	}
	f(doc)
	// fmt.Printf("%v", metas)
	fmt.Println(metas["google-site-verification"])
}
