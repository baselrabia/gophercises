package main

import (
	"bytes"
	"flag"
	"fmt"
	"golang.org/x/net/html"
	"log"
	"os"
	"strings"
)

type Link struct {
	Href string
	Text string
}

func main() {
	flagHtml := flag.String("html", "ex1.html", "the html file that u wanna pass")
	flag.Parse()

	f, err := os.Open(*flagHtml)
	if err != nil {
		log.Fatalf("failed to open the html file %q: %v \n", *flagHtml, err)
	}
	defer f.Close()

	root, err := html.Parse(f)
	if err != nil {
		log.Fatalf("failed to parse html  %v \n", err)
	}

	as := make(chan *html.Node)
	go findAnchors(root, as)
	for a := range as {
		fmt.Println(Link{extractHref(a), extractText(a)})
	}
}

func extractText(a *html.Node) string {
	text := &bytes.Buffer{}
	collectText(a, text)
	return strings.TrimSpace(fmt.Sprintf("%s", text))
}

func collectText(n *html.Node, buf *bytes.Buffer) {
	if n.Type == html.TextNode {
		buf.WriteString(n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		collectText(c, buf)
	}
}

func extractHref(a *html.Node) string {
	for _, attr := range a.Attr {
		if attr.Key == "href" {
			return attr.Val
		}
	}
	return ""
}

func findAnchors(n *html.Node, as chan *html.Node) {
	// the break condition when find an anchor element
	if n.Type == html.ElementNode && n.Data == "a" {
		as <- n
		return
	}

	// moving in the html tree
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		findAnchors(c, as)
	}

	// close the chanel to stop the loop on chan
	if n.Parent == nil {
		close(as)
	}

}
