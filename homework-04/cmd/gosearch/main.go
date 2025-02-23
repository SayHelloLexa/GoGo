package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"go-core-4/homework-02/pkg/crawler"
	"go-core-4/homework-02/pkg/crawler/spider"

	"go-core-4/homework-02/pkg/index"
	"go-core-4/homework-04/pkg/jsonutils"
)

// Функция сканирует страницы по переданным ссылкам и возвращает мапу документов
func scan(idx index.Index, urls ...string) (map[int]crawler.Document, error) {
	c := spider.New()
	storage := make(map[int]crawler.Document)

	// Сканируем страницы 
	for i := range urls {
		var d []crawler.Document

		if jsonutils.IsExist(urls[i]) {
			file, err := os.ReadFile("../../JSON/" + jsonutils.UrlMap(urls[i]) + ".JSON")
			if err != nil {
				return storage, fmt.Errorf("opening file error: %v", err)
			}

			err = json.Unmarshal(file, &d)
			if err != nil {
				return storage, fmt.Errorf("unmarshaling error: %v", err)
			}
		} else {
			var err error
			d, err = c.Scan(urls[i], 2)
			if err != nil {
				return storage, fmt.Errorf("scaning error: %v", err)
			}

			// Создаем директорию и файл для сохранения данных
			filepath, err := jsonutils.CreateDir(urls[i])
			if err != nil {
				return storage, err
			}

			jsonData, err := json.Marshal(d)
			if err != nil {
				return storage, fmt.Errorf("marshaling error: %v", err)
			}

			// Создаем или открываем файл для записи
			file, err := os.Create(filepath)
			if err != nil {
				return storage, fmt.Errorf("opening file error: %v", err)
			}
			defer file.Close()

			_, err = file.Write(jsonData)
			if err != nil {
				return storage, fmt.Errorf("writing file error: %v", err)
			}
		}

		// Добавляем документы в мапу и инвертированный индекс
		for j, v := range d {
			v.ID = j
			idx.Add(v.Title, v.ID)
			storage[j] = v
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
