package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"go-core-4/homework-02/pkg/crawler"
	"go-core-4/homework-02/pkg/crawler/spider"

	"go-core-4/homework-02/pkg/index"
)

// Функция проверяет существование файла
// func isExist(url string) bool {
// 	url = strings.Map(func(r rune) rune {
// 		switch r {
// 		case 'h', 't', 'p', 's', ':', '/':
// 				return -1 // Удалить
// 		case '.':
// 				return '-' // Заменить на '-'
// 		default:
// 				return r // Оставить как есть
// 		}
// 	}, url)

// 	_, err := os.Stat(url)
// 	if err == nil {
// 		return true
// 	}
// 	if os.IsNotExist(err) {
// 		return false
// 	}
// 	return true
// }

// Функция создает директорию для хранения
// txt результатов сканирования страниц
func createDir(url string) (string, error) {
	url = strings.Map(func(r rune) rune {
		switch r {
		case 'h', 't', 'p', 's', ':', '/':
				return -1 // Удалить
		case '.':
				return '-' // Заменить на '-'
		default:
				return r // Оставить как есть
		}
	}, url)

	dir := "../../internal/"
	err := os.MkdirAll(dir, 0777)
	if err != nil {
		return "", fmt.Errorf("creating directory error: %v", err)
	}

	filepath := dir + url + ".JSON"
	file, err := os.Create(filepath)
	if err != nil {
		return "", fmt.Errorf("creating file error: %v", err)
	}
	defer file.Close()

	return filepath, nil
}

// Функция сканирует страницы по переданным ссылкам и возвращает мапу документов
func scan(idx index.Index, urls ...string) (map[int]crawler.Document, error) {
	c := spider.New()
	storage := make(map[int]crawler.Document)

	// Сканируем страницы 
	for i := range urls {
		d, err := c.Scan(urls[i], 2)
		if err != nil {
			return storage, fmt.Errorf("scaning error: %v", err)
		}

		// Создаем директорию и файл для сохранения данных
		filepath, errd := createDir(urls[i])
		if errd != nil {
			return storage, errd
		}

		// Открываем файл для записи
		file, errf := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if errf != nil {
			return storage, fmt.Errorf("opening file error: %v", errf)
		}
		defer file.Close()

		// Создаем JSON-энкодер для записи данных в файл
		encoder := json.NewEncoder(file)

		// Добавляем документы в мапу и инвертированный индекс
		for i, v := range d {
			v.ID = i
			idx.Add(v.Title, v.ID)
			storage[i] = v
			encoder.Encode(v)
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
