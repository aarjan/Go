package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func main() {
	resp, err := http.Get("http://hamrobazaar.com/c10-books-and-learning")

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	root, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Matcher function
	matcher := func(n *html.Node) bool {
		if n.DataAtom == atom.A && n.Parent != nil && n.Parent.Parent != nil && n.Parent.Parent.Parent != nil && n.Parent.Parent.Parent.Parent != nil {
			if n.Parent.DataAtom == atom.Td && scrape.Attr(n.Parent, "height") == "110" {
				return scrape.Attr(n.Parent.Parent.Parent.Parent, "border") == "0" && scrape.Attr(n.Parent.Parent.Parent.Parent, "cellspacing") == "0" && scrape.Attr(n.Parent.Parent.Parent.Parent, "cellpadding") == "0"
			}
		}
		return false
	}

	articles := scrape.FindAll(root, matcher)
	for _, article := range articles {

		fmt.Println(scrape.Text(article.Parent))
		fmt.Println("----------------------")

	}
}
