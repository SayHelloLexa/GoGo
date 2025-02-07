package main

import (
	"flag"
	"fmt"

	"go-core-4/homework-02/pkg/crawler"
	"go-core-4/homework-02/pkg/crawler/spider"

	"go-core-4/homework-02/pkg/index"
)

// Функция сканирует страницы по переданным ссылкам и возвращает мапу документов
func scan(idx index.Index, urls ...string) (map[int]crawler.Document, error) {
	c := spider.New()
	storage := make(map[int]crawler.Document)

	// Сканируем страницы и добавляем документы в мапу
	for i := range urls {
		d, err := c.Scan(urls[i], 2)
		if err != nil {
			return storage, fmt.Errorf("scaning error: %v", err)
		}

		for i, v := range d {
			v.ID = i
			idx.Add(v.Title, v.ID)
			storage[i] = v
		}
	}

	return storage, nil
}

func main() {
	// Обрабатываем флаг
	var f string
	flag.StringVar(&f, "s", "", "search for a word using a link.")
	flag.Parse()

	// Проверяем, что флаг не пустой
	if f == "" {
		flag.PrintDefaults()
		return
	}

	// Создаем инвертированный индекс
	index := index.New()
	
	// Сканируем страницы и получаем мапу документов
	storage, err := scan(index, "https://go.dev", "https://golang.org")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Получаем массив индексов документов, которые соответствуют флагу
	res := index.Search(f)

	// Проходимся по массиву индексов и выводим документы
	for _, v := range res {
		fmt.Printf("Title: %s, ID: %d, URL: %s\n", storage[v].Title, storage[v].ID, storage[v].URL)
	}
}
