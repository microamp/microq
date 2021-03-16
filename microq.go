package microq

import (
	"io"

	"golang.org/x/net/html"
)

type predicate func(*html.Node) bool

func filter(p predicate, n *html.Node) <-chan *html.Node {
	out := make(chan *html.Node)

	go func() {
		defer close(out)

		var fn func(*html.Node)
		fn = func(n *html.Node) {
			if n == nil {
				return
			}
			if p(n) {
				out <- n
			}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				fn(c)
			}
		}
		fn(n)
	}()

	return out
}

func search(ps []predicate, n *html.Node) <-chan *html.Node {
	out := make(chan *html.Node)

	go func() {
		defer close(out)

		var fn func([]predicate, *html.Node)
		fn = func(ps []predicate, n *html.Node) {
			if n == nil {
				return
			}
			if len(ps) == 0 {
				out <- n
				return
			}
			for x := range filter(ps[0], n) {
				fn(ps[1:], x)
			}
		}
		fn(ps, n)
	}()

	return out
}

func Query(reader io.Reader, q string) (<-chan *html.Node, error) {
	root, err := html.Parse(reader)
	if err != nil {
		return nil, err
	}

	ps := predicates(q)
	return search(ps, root), nil
}

func Texts(ns <-chan *html.Node) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)

		for n := range ns {
			var fn func(*html.Node)
			fn = func(n *html.Node) {
				if n.Type == html.TextNode {
					out <- n.Data
				}
				for c := n.FirstChild; c != nil; c = c.NextSibling {
					fn(c)
				}
			}
			fn(n)
		}
	}()

	return out
}
