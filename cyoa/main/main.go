package main

import (
	"flag"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/mroseman95/gophercises/cyoa/story"
)

func main() {
	storyFilePath := flag.String("story", "gopher.json", "filepath to the story JSON file")
	flag.Parse()

	// read in story bytes
	jsonBytes, err := ioutil.ReadFile(*storyFilePath)
	if err != nil && err != io.EOF {
		panic(err)
	}

	storyArcs, err := story.ReadStory(jsonBytes)
	if err != nil {
		panic(err)
	}

	// display intro arc as HTML at localhost:8080
	fmt.Println("Starting the server on :8080")
	http.HandleFunc("/", NewStoryHandler(storyArcs))
	http.ListenAndServe(":8080", nil)
}

// NewStoryHandler creates a new handler func that loads the webpages for the given
// choose your own adventure story.
func NewStoryHandler(arcs story.Arcs) http.HandlerFunc {
	mux := http.NewServeMux()
	for arcTitle, arc := range arcs {
		mux.HandleFunc(fmt.Sprintf("/%s", arcTitle), NewArcHandler(arc))
	}

	return func(w http.ResponseWriter, r *http.Request) {
		mux.ServeHTTP(w, r)
	}
}

// NewArcHandler creates a new handler func that loads a webpage for a given arc
func NewArcHandler(arc story.Arc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO hand option links, possibly go to another endpoint on the server
		t, err := template.ParseFiles("index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		t.Execute(w, arc)
	}
}
