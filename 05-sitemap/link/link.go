package link

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"strings"
)

type Link struct {
	Href string
	Text string
}

func Parse(reader io.Reader) ([]Link, error) {

	root, err := html.Parse(reader)
	if err != nil {
		return nil, err
	}

	as := make(chan *html.Node)
	go findAnchors(root, as)
	var links []Link
	for a := range as {
		link := Link{extractHref(a), extractText(a)}
		links = append(links, link)
	}

	return links, nil
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
