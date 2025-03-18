package main

import (
	"io"
	"os"
)

func GetString(w io.Writer, a ...any) {
	for _, v := range a {
		if v, ok := v.(string); ok {
			w.Write([]byte(v))
		}
	}
}

func main() {
	GetString(os.Stdout, "Hello", 123, "World")
}