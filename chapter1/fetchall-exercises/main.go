// Fetchall exercises
// Exercise 1.10 - Check if time changes with a website with large amount of data
// print output of fetchall to a file
// Exercise 1.11 - Use a list of top million web sites available at alexa.com
// and see how program behaves
package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, urlvar := range os.Args[1:] {
		// Adding prefix if it is missing
		if !strings.HasPrefix(urlvar, "http://") {
			urlvar = "http://" + urlvar
		}
		go fetch(urlvar, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) //receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(urlvar string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(urlvar)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	u, err := url.Parse(urlvar)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't parse url: %s\n", err)
		os.Exit(1)
	}
	// This to extract just the name of the host
	f, err := os.Create(u.Host + ".txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't create file for %s\n", urlvar)
		os.Exit(1)
	}
	// So File returned by os.Create is a Writer also
	nbytes, err := io.Copy(f, resp.Body)
	resp.Body.Close() // don't leak resources
	f.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", urlvar, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d  %s", secs, nbytes, urlvar)
}
