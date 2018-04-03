package story

import (
	"bytes"
	"encoding/json"
	"io"
)

// Arc describes a chapter/arc of the story
type Arc struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

// Option describes the various options you can follow to continue the story
type Option struct {
	Text    string `json:"text"`
	NextArc string `json:"arc"`
}

// Arcs is a map of each arc name to the Arc struct
type Arcs map[string]Arc

// ReadStory takes some json in bytes, and reads in the story, parsing the bytes, and
// returns a map of arc names to arc structs
func ReadStory(jsonBytes []byte) (Arcs, error) {
	jsonReader := bytes.NewReader(jsonBytes)
	dec := json.NewDecoder(jsonReader)
	arcs := make(Arcs)
	for {
		err := dec.Decode(&arcs)
		if err == io.EOF {
			return arcs, nil
		}
		if err != nil {
			return nil, err
		}
	}
}
