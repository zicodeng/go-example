# Tokenizing HTML Streams

The code creates a program that will fetch some HTML and extract the value of its `<title>` element.

## Installation

Install all dependencies.

    go get

## Usage

Move to extract-title.go directory.

    cd path/to/extract-title.go

Get page title for one URL.

    go run extract-title.go http://example.com

Get page titles for multiple URLs.

    go run extract-title.go http://example.com http://google.com

## Reference

https://drstearns.github.io/tutorials/tokenizing/