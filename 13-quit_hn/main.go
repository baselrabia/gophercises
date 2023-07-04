package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"quite_hn/hn"
	"sort"
	"strings"
	"sync"
	"time"
)

var (
	cache     = map[int]hn.Item{}
	cacheLock sync.RWMutex
)

func main() {
	// parse flags
	var port, numStories int
	flag.IntVar(&port, "port", 3000, "the port to start the web server on")
	flag.IntVar(&numStories, "num_stories", 30, "the number of top stories to display")
	flag.Parse()

	tpl := template.Must(template.ParseFiles("./index.gohtml"))

	http.HandleFunc("/", handler(numStories, tpl))

	// Start the server
	fmt.Println("Listen And Serve on: http://127.0.0.1:3000")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func handler(numStories int, tpl *template.Template) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		var client hn.Client
		ids, err := client.TopItems()
		if err != nil {
			http.Error(w, "Failed to load top stories", http.StatusInternalServerError)
			return
		}

		ids = ids[:int(float32(numStories)*1.25)]
		storyChan := make(chan orderItem)
		var wg sync.WaitGroup
		for i, id := range ids {
			wg.Add(1)
			go func(id, idx int) {
				defer wg.Done()

				if _, ok := cacheRead(id); !ok {
					item, err := client.GetItem(id)
					if err != nil {
						return
					}
					cacheWrite(id, item)
				}
				hnItem, _ := cacheRead(id)
				item := parseHNItem(hnItem)
				if isStoryLink(item) {
					storyChan <- orderItem{item, idx}
				}
			}(id, i)
		}

		go func() {
			wg.Wait()
			close(storyChan)
		}()

		var stories []orderItem

		for orderItem := range storyChan {
			stories = append(stories, orderItem)
			if len(stories) >= numStories {
				break
			}
		}
		sort.Slice(stories, func(i, j int) bool {
			return stories[i].idx < stories[j].idx
		})

		data := templateData{
			Stories: stories,
			Time:    time.Now().Sub(start),
		}
		err = tpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Failed to process the template", http.StatusInternalServerError)
			return
		}
	})
}

func cacheRead(id int) (hn.Item, bool) {
	cacheLock.RLock()
	defer cacheLock.RUnlock()
	item, ok := cache[id]
	return item, ok
}

func cacheWrite(id int, item hn.Item) {
	cacheLock.Lock()
	defer cacheLock.Unlock()
	cache[id] = item
}

func isStoryLink(item item) bool {
	return item.Type == "story" && item.URL != ""
}

func parseHNItem(hnItem hn.Item) item {
	ret := item{Item: hnItem}
	url, err := url.Parse(ret.URL)
	if err == nil {
		ret.Host = strings.TrimPrefix(url.Hostname(), "www.")
	}
	return ret
}

// item is the same as the hn.Item, but adds the Host field
type item struct {
	hn.Item
	Host string
}

type orderItem struct {
	item
	idx int
}
type templateData struct {
	Stories []orderItem
	Time    time.Duration
}
