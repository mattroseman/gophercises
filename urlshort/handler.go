package urlshort

import (
	"fmt"
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

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
		} else {
			// if the url path isn't in the pathsToUrls map, use the fallback handler
			fallback.ServeHTTP(w, r)
		}
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
	// parse yml into a slice of redirect types
	var redirects []redirect
	err := yaml.Unmarshal(yml, &redirects)
	if err != nil {
		return nil, err
	}

	// turn the slice of redirect types into a map
	ymlMap := make(map[string]string)
	for _, r := range redirects {
		ymlMap[r.Path] = r.URL
	}

	// call MapHandler with this new map made from the YAML
	return MapHandler(ymlMap, fallback), nil
}

type redirect struct {
	Path string
	URL  string
}
