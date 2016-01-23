package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/url", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		resp, err := http.Get("https://github.com/")
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		doc, err := html.Parse(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		metamap := MetaParser(doc)
		result, err := json.Marshal(metamap)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.printf("%s\n", result)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "%s\n", result)
	})

	http.Handle("/", http.FileServer(http.Dir("./public")))

	log.Fatal(http.ListenAndServe(":4000", nil))
}
