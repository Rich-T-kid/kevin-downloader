package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

const source = "playlist_content.html"
const matchTerm = "href"

func ExtractHref(line []byte) string {
	// Find 'href="'
	hrefStart := bytes.Index(line, []byte(`href="`))
	if hrefStart == -1 {
		return ""
	}

	// Move past 'href="'
	urlStart := hrefStart + 6

	// Find the closing quote
	urlEnd := bytes.Index(line[urlStart:], []byte(`"`))
	if urlEnd == -1 {
		return ""
	}

	// Extract the URL
	return string(line[urlStart : urlStart+urlEnd])
}
func decodeHTML(file_source string) []string {
	f, err := os.Open(file_source)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	content, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	var allLinks []string
	for _, nlChunk := range bytes.Split(content, []byte("\n")) {
		if bytes.Contains(nlChunk, []byte(matchTerm)) {
			link := ExtractHref(nlChunk)
			if link != "" { // in case of parsing issues
				allLinks = append(allLinks, link)
			}
		}
	}
	outputF, err := os.OpenFile("links.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	for _, l := range allLinks {
		outputF.Write([]byte(l))
		outputF.Write([]byte("\n"))
	}
	return allLinks
}

func main() {
	links := decodeHTML(source)
	for i, link := range links {
		if i > 2 {
			break
		}
		resp, err := http.Get(link)
		if err != nil {
			fmt.Printf("error occured making get request to %v: %v", link, err)
		}
		content, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("error occured making get request to %v: %v", link, err)
		}
		outputF, err := os.OpenFile(fmt.Sprintf("link-%d", i), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			continue // skip
		}
		outputF.Write(content)

	}
}
