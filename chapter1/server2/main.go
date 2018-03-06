// Server2
package main

import (
	"fmt"
	"log"
	"net/http"
)

var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", countHandler)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/count" {
		count++
	}
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func countHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Request count: %d\n", count)
}
