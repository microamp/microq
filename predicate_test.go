package microq

import (
	"testing"
)

func TestTokenHasClass(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{
			input:    "p",
			expected: false,
		},
		{
			input:    "p#id",
			expected: false,
		},
		{
			input:    "p.class",
			expected: true,
		},
	}
	for _, test := range tests {
		tk := token(test.input)
		received := tk.hasClass()
		if received != test.expected {
			t.Errorf("Received %t, expected %t", received, test.expected)
		}
	}
}

func TestTokenGetClassVal(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "p.hello",
			expected: "hello",
		},
		{
			input:    "p.world",
			expected: "world",
		},
	}
	for _, test := range tests {
		tk := token(test.input)
		received := tk.getClassVal()
		if received != test.expected {
			t.Errorf("Received %s, expected %s", received, test.expected)
		}
	}
}

func TestTokenGetData(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "p",
			expected: "p",
		},
		{
			input:    "p.class",
			expected: "p",
		},
		{
			input:    "p#id",
			expected: "p",
		},
	}
	for _, test := range tests {
		tk := token(test.input)
		received := tk.getData()
		if received != test.expected {
			t.Errorf("Received %s, expected %s", received, test.expected)
		}
	}
}

func TestTokenise(t *testing.T) {
	input := "ol.repo-list li.d-block"
	tokens := tokenise(input)

	if len(tokens) != 2 {
		t.Errorf("Length expected %d, %d received", 2, len(tokens))
	}
	if tokens[0] != "ol.repo-list" {
		t.Errorf("Received %s, expected %s", tokens[0], "ol.repo-list")
	}
	if tokens[1] != "li.d-block" {
		t.Errorf("Received %s, expected %s", tokens[1], "li.d-block")
	}
}
