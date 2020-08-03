package engine

import "golang.org/x/net/html"

type Request struct {
	Url       string
	Parsefunc func(node *html.Node) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}
