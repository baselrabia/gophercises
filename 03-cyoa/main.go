package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type Story map[string]struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	}
}

func main() {
	flagStoryJSONFilename := flag.String("story", "gopher.json", "the json file that contain the story")
	flag.Parse()

	f, err := os.Open(*flagStoryJSONFilename)
	if err != nil {
		fmt.Printf("failed to open file %s: %v", *flagStoryJSONFilename, err)
		return
	}
	defer f.Close()

	var story Story
	if err := json.NewDecoder(f).Decode(&story); err != nil {
		fmt.Printf("failed to decode file %v", err)
		return
	}

	fmt.Printf("server is up and running 8080...\n")
	mux := newStoryMux(story)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}

type StoryMux struct {
	story Story
}

func newStoryMux(story Story) http.Handler {
	return &StoryMux{story}
}

func (m *StoryMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	arcName := r.URL.Query().Get("arc")
	if arcName == "" {
		arcName = "intro"
	}

	arc, ok := m.story[arcName]
	if !ok {
		http.Error(w, fmt.Sprintf("arc not found: %s", arc), http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("arc.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to parse template: %v", err), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, arc); err != nil {
		http.Error(w, fmt.Sprintf("failed to execute template: %v", err), http.StatusInternalServerError)
		return
	}

}
