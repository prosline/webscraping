package link

import (
	"io"
	"strings"
	"golang.org/x/net/html"
)
type Link struct{
	Href string
	Text string
}

// Parse html tags for href tags
func Parse(r io.Reader)([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil{
		return nil,err
	}
	nodes := linkNodes(doc)
	var links []Link
	for _, node := range nodes {
		links = append(links, linkBuilder(node))
	}
	return links, nil
}
// Retrieve the text from a Href tag
func getText(t *html.Node) string {
	if t.Type == html.TextNode {
		return t.Data
	}
	if t.Type != html.ElementNode {
		return ""
	}
	var ret string
	for c := t.FirstChild; c != nil; c = c.NextSibling {
		ret += getText(c) + " "
	} 
	fields := strings.Fields(ret)
	ret = strings.Join(fields," ")
	return ret
}
// Build a Link object based on 'href tag` fund
func linkBuilder(n *html.Node) Link {
	var ret  Link
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
			break
		}
	}
	ret.Text = getText(n)
	return ret
}

// Parse html tags and return a slice of html.Node object
// containing href tags
func linkNodes(n *html.Node)([]*html.Node){
	if n.Type == html.ElementNode && n.Data == "a"{
		return []*html.Node{n}
	}
	var ret []*html.Node
	for c :=  n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret,linkNodes(c)...)
	}
	return ret
}
