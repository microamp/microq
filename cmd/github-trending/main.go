package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/microamp/microq"
)

func fetch(url string) (io.Reader, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return strings.NewReader(string(body)), nil
}

func main() {
	url := "https://github.com/trending"
	q := "article.Box-row h1.h3 a"

	fmt.Println("Fetching trending repositories from GitHub...")

	r, err := fetch(url)
	if err != nil {
		panic(err)
	}

	ns, err := microq.Query(r, q)
	if err != nil {
		panic(err)
	}

	rank := 1
	for n := range ns {
		for _, a := range n.Attr {
			if a.Key == "href" {
				fmt.Printf("#%02d: https://github.com/%s\n", rank, a.Val)
				break
			}
		}
		rank++
	}
}
