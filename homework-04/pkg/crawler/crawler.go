package crawler

// Поисковый робот.
// Осуществляет сканирование сайтов.

// Interface определяет контракт поискового робота.
type Interface interface {
	Scan(url string, depth int) ([]Document, error)
	BatchScan(urls []string, depth int, workers int) (<-chan Document, <-chan error)
}

// Document - документ, веб-страница, полученная поисковым роботом.
type Document struct {
    ID    int    `json:"ID"`
    URL   string `json:"URL"`
    Title string `json:"Title"`
    Body  string `json:"Body"`
}
