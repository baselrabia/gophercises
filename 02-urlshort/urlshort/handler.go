package urlshort

import (
	"encoding/json"
	"fmt"
	"github.com/go-yaml/yaml"
	"net/http"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		path := request.URL.Path
		longUrlPath, ok := pathsToUrls[path]
		if !ok {
			fallback.ServeHTTP(writer, request)
		}
		//http.Redirect(writer, request, longUrlPath, http.StatusPermanentRedirect)
		fmt.Fprintf(writer, longUrlPath)
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
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.

type data struct {
	Path string
	Url  string
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var parsedYamlData []data
	if err := yaml.Unmarshal([]byte(yml), &parsedYamlData); err != nil {
		return nil, err
	}
	pathsToUrls := map[string]string{}
	for _, entity := range parsedYamlData {
		pathsToUrls[entity.Path] = entity.Url
	}

	return MapHandler(pathsToUrls, fallback), nil
}

func JSONHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var parsedJsonData []data
	if err := json.Unmarshal([]byte(yml), &parsedJsonData); err != nil {
		return nil, err
	}
	pathsToUrls := map[string]string{}
	for _, entity := range parsedJsonData {
		pathsToUrls[entity.Path] = entity.Url
	}

	return MapHandler(pathsToUrls, fallback), nil
}
