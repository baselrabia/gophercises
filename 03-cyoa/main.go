package main

import (
	"encoding/json"
	"flag"
	"fmt"
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

	fmt.Println(story)

}
