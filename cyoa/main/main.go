package main

import (
	"flag"
	"io"
	"io/ioutil"

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

	_, err = story.ReadStory(jsonBytes)
	if err != nil {
		panic(err)
	}
}
