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
	// metas := map[string]string{}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "meta" {
			fmt.Println(n.Attr)
			//for _, a := range n.Attr {
			//fmt.Println(a.Key)
			// if a.Key == "name" {
			// 	fmt.Println("name: ", a.Val)
			// 	break
			// }
			// if a.Key == "content" {
			// 	fmt.Println("content: ", a.Val)
			// 	break
			// }
			//}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
		// fmt.Printf("%v", metas)
	}
	f(doc)
}
