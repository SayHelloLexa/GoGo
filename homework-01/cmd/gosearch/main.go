package main

import (
	"flag"
	"fmt"
	"strings"

	"go-core-4/homework-01/pkg/crawler"
	"go-core-4/homework-01/pkg/crawler/spider"
)

func scan(urls ...string) (map[string]crawler.Document, error) {
	c := spider.New()
	storage := make(map[string]crawler.Document)

	for i := range urls {
		d, err := c.Scan(urls[i], 2)
		if err != nil {
			return storage, fmt.Errorf("scaning error: %v", err)
		}

		for _, v := range d {
			storage[v.Title] = v
		}
	}

	return storage, nil
}

func main() {
	var f string
	flag.StringVar(&f, "s", "", "search for a word using a link")
	flag.Parse()

	urlStorage := make(map[string]string)

	storage, err := scan("https://go.dev", "https://golang.org")
	if err != nil {
		fmt.Println(err)
		return
	}

	if f == "" {
		for k, v := range storage {
			s := fmt.Sprintf("Storage key: %s\nID: %v\nURL: %s\nTitle: %s\n \n", k, v.ID, v.URL, v.Title)
			
			fmt.Print(s)
		}
	} else {
		for _, v := range storage {
			if strings.Contains(strings.ToLower(v.Title), strings.ToLower(f)) {
				urlStorage[v.Title] = v.URL
			}
		}

		if len(urlStorage) == 0 {
			fmt.Println("No matches found")
		} else {
			for k, v := range urlStorage {
				fmt.Printf("Title: %s\nURL: %s\n\n", k, v)
			}
		}
	}
}