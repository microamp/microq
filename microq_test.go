package microq

import (
	"testing"

	"golang.org/x/net/html"
)

func TestFilterEmptyNode(t *testing.T) {
	test := struct {
		p predicate
		n *html.Node
	}{p: nil, n: nil}

	var ns []*html.Node
	for n := range filter(test.p, test.n) {
		ns = append(ns, n)
	}
	if len(ns) > 0 {
		t.Errorf("Non-empty channel (size: %d)", len(ns))
	}
}

func TestSearchNoPredicates(t *testing.T) {
	test := struct {
		ps []predicate
		n  *html.Node
	}{
		ps: nil, n: &html.Node{Type: html.TextNode, Data: "hello"},
	}

	var ns []*html.Node
	for n := range search(test.ps, test.n) {
		ns = append(ns, n)
	}
	if len(ns) != 1 {
		t.Fatalf("Channel sizes differ: %d vs %d", len(ns), 1)
	}
	if ns[0].Type != html.TextNode {
		t.Errorf("Node types differ: %v vs %v", ns[0].Type, html.TextNode)
	}
	if ns[0].Data != "hello" {
		t.Errorf("Node data differ: %s vs %s", ns[0].Data, "hello")
	}
}

func TestSearchEmptyNode(t *testing.T) {
	test := struct {
		ps []predicate
		n  *html.Node
	}{
		ps: nil, n: nil,
	}

	var ns []*html.Node
	for n := range search(test.ps, test.n) {
		ns = append(ns, n)
	}
	if len(ns) > 0 {
		t.Error(ns[0])
		t.Errorf("Non-empty channel (size: %d)", len(ns))
	}
}
