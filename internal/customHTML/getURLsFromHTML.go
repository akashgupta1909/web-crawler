package customHTML

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func GetURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	parsedHTML, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error in parsing html: %v", err))
	}

	urls := []string{}
	var traverseNode func(node *html.Node)

	traverseNode = func(node *html.Node) {
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			if child.Type == html.ElementNode && child.DataAtom == atom.A {
				for _, attr := range child.Attr {
					if attr.Key == "href" {
						href, err := url.Parse(attr.Val)
						if err != nil {
							fmt.Printf("error in parsing href: %v", err)
							continue
						}
						urls = append(urls, baseURL.ResolveReference(href).String())
					}
				}
			}
			traverseNode(child)
		}
	}
	traverseNode(parsedHTML)
	return urls, nil
}
