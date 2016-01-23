package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/fetch", func(resp http.ResponseWriter, req *http.Request) {
		// fmt.Fprintf(resp, "Hello, %q", html.EscapeString(req.URL.Path))

		resp.Header().Set("Content-Type", "application/json")

		getresp, err := http.Get(req.FormValue("url"))

		if err != nil {
			// log.Fatal(err)
			resp.Header().Set("Status", "500 Internal Server Error")
			message := map[string]string{}
			message["message"] = "Internal Server Error"
			emessage, err := json.MarshalIndent(message, "", "\t")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Fprintf(resp, "%s\n", emessage)
			return
		} else {
			defer getresp.Body.Close()
		}

		doc, err := html.Parse(getresp.Body)
		if err != nil {
			log.Fatal(err)
		}
		metamap := MetaParser(doc)
		result, err := json.MarshalIndent(metamap, "", "\t")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintf(resp, "%s\n", result)
	})

	http.Handle("/", http.FileServer(http.Dir("./public")))

	log.Fatal(http.ListenAndServe(":4000", nil))
}
