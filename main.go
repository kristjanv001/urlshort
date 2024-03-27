package main

import (
	"fmt"
	"net/http"
	"urlshortener/handler"
)

func main() {
	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/wikipedia": "https://wikipedia.com",
		"/go":     "https://go.dev/",
	}
	mapHandler := handler.MapHandler(pathsToUrls, mux)

	yaml := `
- path: /google
  url: https://google.com
- path: /x
  url: https://x.com
`
	yamlHandler, err := handler.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
