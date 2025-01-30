package main

import (
	"fmt"

	"go-core-4/homework-01/pkg/crawler"
	"go-core-4/homework-01/pkg/crawler/spider"
)

func main() {
	c := spider.New()
	storage := make(map[string]crawler.Document)

	d, err := c.Scan("https://go.dev", 2)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	for _, v := range d {
		storage[v.Title] = v
	}

	fmt.Println(storage)
}