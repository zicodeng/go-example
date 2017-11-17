package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

const headerContentType = "Content-Type"
const contentTypeHTML = "text/html"

// PageLinks contains summary information about a web page.
type PageLinks struct {
	Title string   `json:"title"`
	Links []string `json:"links"`
}

// GetPageLinks fetches PageLinks info for a given URL.
func GetPageLinks(URL string) (*PageLinks, error) {
	// Parse the URL to get a base URL for relative links.
	baseURL, err := url.Parse(URL)
	if err != nil {
		return nil, fmt.Errorf("error parsing base URL: %v", err)
	}
	// Fetch the URL.
	resp, err := http.Get(URL)
	if err != nil {
		return nil, fmt.Errorf("error getting URL %s: %v", URL, err)
	}
	defer resp.Body.Close()

	// If not OK, return an error.
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error response status code %d while fetching %s", resp.StatusCode, URL)
	}

	// If the requested URL is not an HTML page,
	// just return an empty PageLinks structure.
	if !strings.HasPrefix(resp.Header.Get(headerContentType), contentTypeHTML) {
		return &PageLinks{}, nil
	}

	links := &PageLinks{}
	tokenizer := html.NewTokenizer(resp.Body)
	for {
		ttype := tokenizer.Next()
		if ttype == html.ErrorToken {
			err := tokenizer.Err()
			// If we reached the end of the stream
			// return the PageLinks.
			if err == io.EOF {
				return links, nil
			}
			return links, err
		}

		// If this is a start tag token
		if ttype == html.StartTagToken {
			token := tokenizer.Token()
			// If this is the page title
			if token.Data == "title" {
				tokenizer.Next()
				links.Title = tokenizer.Token().Data
			}

			// If this is a hyperlink...
			if token.Data == "a" {
				// Get the href attribute.
				for _, attr := range token.Attr {
					// Ignore bookmark links.
					if attr.Key == "href" && !strings.HasPrefix(attr.Val, "#") {
						// Parse the link and if there's an error (bad URL)
						// just ignore it.
						link, err := url.Parse(attr.Val)
						if err != nil {
							continue
						}
						// If the link is not absolute
						// make it absolute using the baseURL.
						if !link.IsAbs() {
							link = baseURL.ResolveReference(link)
						}
						links.Links = append(links.Links, link.String())
					}
				} // for all attributes
			} // if <a>
		} // if start tag
	} // for each token
} // GetPageSummary()
