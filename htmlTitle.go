package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func fetchHtmlTitle(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return "failed to GET"
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return "failed to get body"
	}

	html_src := string(body)
	titleStart := strings.Index(html_src, "<title>") + len("<title>")
	titleEnd := strings.Index(html_src, "</title>")
	if titleStart == -1 || titleEnd == -1 || titleEnd < titleStart {
		return "failed to find title"
	}
	return html.UnescapeString(html_src[titleStart:titleEnd])
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: webTitle <url>\n")
		os.Exit(1)
	}

	url := os.Args[1]
	fmt.Printf(fetchHtmlTitle(url) + "\n")
}
