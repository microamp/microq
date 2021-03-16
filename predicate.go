package microq

import (
	"strings"

	"golang.org/x/net/html"
)

const (
	symbolClass = "."
	symbolID    = "#"
)

type token string

func (t token) hasSymbol(symbol string) bool {
	return strings.Contains(string(t), symbol)
}

func (t token) getAttrVal(symbol string) string {
	return strings.Split(string(t), symbol)[1]
}

func (t token) hasClass() bool {
	return t.hasSymbol(symbolClass)
}

func (t token) getClassVal() string {
	return t.getAttrVal(symbolClass)
}

func (t token) hasID() bool {
	return t.hasSymbol(symbolID)
}

func (t token) getIDVal() string {
	return t.getAttrVal(symbolID)
}

func (t token) getData() string {
	if t.hasClass() {
		return strings.Split(string(t), symbolClass)[0]
	}
	if t.hasID() {
		return strings.Split(string(t), symbolID)[0]
	}
	return string(t)
}

func (token) matchingAttrVal(n *html.Node, key, val string) bool {
	for _, a := range n.Attr {
		if a.Key == key {
			for _, v := range strings.Split(a.Val, " ") {
				if v == val {
					return true
				}
			}
		}
	}
	return false
}

func pred(t token) predicate {
	return func(n *html.Node) bool {
		if n.Type != html.ElementNode {
			return false
		}
		if n.Data != t.getData() {
			return false
		}
		if t.hasClass() && !t.matchingAttrVal(n, "class", t.getClassVal()) {
			return false
		}
		if t.hasID() && !t.matchingAttrVal(n, "id", t.getIDVal()) {
			return false
		}
		return true
	}
}

func tokenise(q string) []token {
	var ts []token
	for _, s := range strings.Split(q, " ") {
		ts = append(ts, token(s))
	}
	return ts
}

func predicates(q string) []predicate {
	var ps []predicate
	for _, t := range tokenise(q) {
		ps = append(ps, pred(t))
	}
	return ps
}
