package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"urlshort/urlshort"
)

func main() {
	flagYamlFilename := flag.String("yml", "urls.yaml", "Path to yaml file containing path/url mappings")
	flagJsonFilename := flag.String("json", "urls.json", "Path to json file containing path/url mappings")
	flag.Parse()

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	//	yaml := `
	//- path: /urlshort
	//  url: https://github.com/gophercises/urlshort
	//- path: /urlshort-final
	//  url: https://github.com/gophercises/urlshort/tree/solution
	//`

	ymlfile, err := os.Open(*flagYamlFilename)
	if err != nil {
		fmt.Printf("failed to open yaml file %q : %v\n", *flagYamlFilename, err)
		return
	}
	defer ymlfile.Close()

	yaml, err := io.ReadAll(ymlfile)
	if err != nil {
		fmt.Printf("failed to read yaml file %q : %v\n", *flagYamlFilename, err)
		return
	}
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}

	/// Json

	jsonfile, err := os.Open(*flagJsonFilename)
	if err != nil {
		fmt.Printf("failed to open yaml file %q : %v\n", *flagJsonFilename, err)
		return
	}
	defer jsonfile.Close()

	json, err := io.ReadAll(jsonfile)
	if err != nil {
		fmt.Printf("failed to read yaml file %q : %v\n", *flagJsonFilename, err)
		return
	}
	jsonHandler, err := urlshort.JSONHandler([]byte(json), mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", jsonHandler)
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
