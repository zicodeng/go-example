package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	// os.Args allows us to access the commands typed by the user on the shell.
	if len(os.Args) < 2 {
		// Print the usage and exit with an error.
		fmt.Println("Usage: go run extract-title.go <url> <url> ...")
		// 0 indicates success, 1 error.
		os.Exit(1)
	}

	// os.Args[0] is the commands used to execute the program.
	// os.Args[1] is the first parameter.
	for index, URL := range os.Args {
		if index > 0 {
			title, err := fetchTitle(URL)
			// If an error occurs, report it and exit.
			if err != nil {
				log.Fatalf("Error fetching page title: %v\n", err)
			}

			// Print the page title.
			fmt.Println(title)
		}
	}
}

// fetchHTML fetches the provided URL and returns the response body or an error.
func fetchHTML(URL string) (io.ReadCloser, error) {
	// Fetch the URL.
	// http.Get() function returns a pointer to an http.Response struct and potentially an error.
	res, err := http.Get(URL)
	if err != nil {
		return res.Body, fmt.Errorf("fetching URL failed %v", err)
	}

	// Verify response status code.
	if res.StatusCode != http.StatusOK {
		return res.Body, fmt.Errorf("response status code was %d", res.StatusCode)
	}

	// Verify response content type
	contentType := res.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "text/html") {
		return res.Body, fmt.Errorf("response content type was %s, not text/html", contentType)
	}

	// Return without error.
	return res.Body, nil
}

//extractTitle returns the content within the <title> element or an error.
func extractTitle(body io.ReadCloser) (string, error) {

	// Create a new tokenizer over the response body.
	tokenizer := html.NewTokenizer(body)

	// Loop until we find the title element and its content
	// or encounter an error (which includes the end of the stream).
	for {
		// Get the next token type.
		tokenType := tokenizer.Next()

		// If it is an error token, we either reached
		// the end of the file, or the HTML was malformed.
		if tokenType == html.ErrorToken {
			return "", fmt.Errorf("error tokenizing HTML: %v", tokenizer.Err())
		}

		// If this is a HTML start-tag token...
		if tokenType == html.StartTagToken {

			// Get the start-tag's tag name.
			// If it returns "title", it means we find <title> tag.
			if "title" == tokenizer.Token().Data {
				// Get the next token type which should be text with the page's title.
				tokenType = tokenizer.Next()

				// Just make sure it is actually a text token.
				if tokenType == html.TextToken {
					// Return the page title
					// and break out of the loop.
					return tokenizer.Token().Data, nil
				}
			}
		}
	}
}

//fetchTitle fetches the page title for a URL.
func fetchTitle(URL string) (string, error) {
	body, err := fetchHTML(URL)
	if err != nil {
		return "", fmt.Errorf("fetching URL failed: %v", err)
	}

	// Close the response body after we are done with it.
	// The statement is executed at the end of its enclosing function (fetchTitle).
	defer body.Close()

	title, err := extractTitle(body)
	if err != nil {
		return "", fmt.Errorf("extracting title failed: %v", err)
	}

	return title, nil
}
