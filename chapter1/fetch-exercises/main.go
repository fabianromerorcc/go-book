// Fetch exercise 1.7
// I'm replacing ioutil.ReadAll with io.Copy, to copy the response body to os.Stdout
// without requiring a buffer large enough to hold the entire stream
// Exercise 1.8 - Add http:// prefix if is missing
// Exercise 1.9 - Print HTTP status code
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		// Adding prefix if it is missing
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// Because I'm omiting the first variable (bytes written) I can't use short variable initialization
		// as err is not a new variable. If I'd be using a variable for bytes written, then I can use ':='
		fmt.Printf("HTTP status code: %v\n", resp.Status)
		_, err = io.Copy(os.Stdout, resp.Body)
		//resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		//fmt.Printf("%s", b)
	}
}
