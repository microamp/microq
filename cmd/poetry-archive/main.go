package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

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
	urls := make(chan string)
	go func() {
		defer close(urls)

		for page := 1; page < 59; page++ {
			// Type: POEMS, Theme: EMOTIONS & FEELINGS
			url := fmt.Sprintf(
				"https://poetryarchive.org/explore/page/%d/?key&type=poems&theme=394&form&region#038;type=poems&theme=394&form&region",
				page,
			)
			r, err := fetch(url)
			if err != nil {
				panic(err)
			}

			ns, err := microq.Query(r, "div#explore-content div.poem-title h3 a")
			if err != nil {
				panic(err)
			}

			for n := range ns {
				for _, a := range n.Attr {
					if a.Key == "href" {
						urls <- a.Val
						break
					}
				}
			}
		}
	}()

	texts := make(chan string)
	go func() {
		defer close(texts)

		for url := range urls {
			r, err := fetch(url)
			if err != nil {
				panic(err)
			}

			ns, err := microq.Query(r, "div.poem-content")
			if err != nil {
				panic(err)
			}

			for t := range microq.Texts(ns) {
				t = strings.TrimSpace(t)
				t = strings.TrimRight(t, "_")
				t = strings.ReplaceAll(t, "—", "-")
				t = strings.ReplaceAll(t, "‘", "'")
				t = strings.ReplaceAll(t, "’", "'")
				t = strings.ReplaceAll(t, "“", `"`)
				t = strings.ReplaceAll(t, "”", `"`)
				t = strings.ReplaceAll(t, "…", "...")
				if t != "" {
					texts <- t
				}
			}

			time.Sleep(10 * time.Second)
		}
	}()

	for t := range texts {
		fmt.Println(t)
	}
}
