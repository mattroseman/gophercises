package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/mroseman95/gophercises/urlshort"
)

func main() {
	ymlFilePath := flag.String("yml", "redirects.yml", "filepath to the redirects YAML file")
	jsonFilePath := flag.String("json", "redirects.json", "filepath to the redirects JSON file")
	flag.Parse()

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	yaml, err := ioutil.ReadFile(*ymlFilePath)
	if err != nil && err != io.EOF {
		panic(err)
	}

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yamlHandler, err := urlshort.YAMLHandler(yaml, mapHandler)
	if err != nil {
		panic(err)
	}

	json, err := ioutil.ReadFile(*jsonFilePath)
	if err != nil && err != io.EOF {
		panic(err)
	}

	// Build the JSONHandler using the YAMLHandler as the
	// fallback
	jsonHandler, err := urlshort.JSONHandler(json, yamlHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", jsonHandler)
}

func defaultMux() *http.ServeMux {
	// ServeMux maps url patterns to handler functions
	// Here mux maps the '/' pattern to the hello function
	// '/' patter will match every incoming http request
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
