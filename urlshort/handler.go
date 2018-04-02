package urlshort

import (
	"encoding/json"
	"fmt"
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

// redirect represents a URL redirect. Path is the input and it maps to URL
type redirect struct {
	Path string
	URL  string
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if redirect, ok := pathsToUrls[r.URL.Path]; ok {
			// if this url path is a key in pathsToUrls, redirect to its value
			fmt.Printf("%s redirected to %s\n", r.URL, redirect)
			http.Redirect(w, r, redirect, 303)
			return
		}

		// if the url path isn't in the pathsToUrls map, use the fallback handler
		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	ymlMap, err := redirectYAMLToMap(yml)
	if err != nil {
		return nil, err
	}

	// call MapHandler with this new map made from the YAML
	return MapHandler(ymlMap, fallback), nil
}

// redirectYAMLToMap takes a slice of bytes representing a YAML document, and converts it to
// a map of path strings to the url redirect string
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
func redirectYAMLToMap(yml []byte) (ymlMap map[string]string, err error) {
	// parse yml into a slice of redirect types
	var redirects []redirect
	err = yaml.Unmarshal(yml, &redirects)
	if err != nil {
		return
	}

	ymlMap = make(map[string]string)
	// turn the slice of redirect types into a map
	for _, r := range redirects {
		ymlMap[r.Path] = r.URL
	}

	return
}

// JSONHandler will parse the provided JSON and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the JSON, then the
// fallback http.Handler will be called instead.
//
// JSON is expected to be in the format:
//
//		[
//			{path: /some-path, url: https://www.some-url.com/demo},
//			...
//		]
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func JSONHandler(jsonBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	jsonMap, err := redirectJSONToMap(jsonBytes)
	if err != nil {
		return nil, err
	}

	// call MapHandler with this new map made from the YAML
	return MapHandler(jsonMap, fallback), nil
}

// redirectJSONToMap takes a slice of bytes representing a JSON document, and converts it to
// a map of path strings to the url redirect string
// JSON is expected to be in the format:
//
//		[
//			{path: /some-path, url: https://www.some-url.com/demo},
//			...
//		]
func redirectJSONToMap(jsonBytes []byte) (jsonMap map[string]string, err error) {
	var redirects []redirect
	err = json.Unmarshal(jsonBytes, &redirects)
	if err != nil {
		return
	}

	jsonMap = make(map[string]string)
	// turn the slice of redirect types into a map
	for _, r := range redirects {
		jsonMap[r.Path] = r.URL
	}

	return
}
