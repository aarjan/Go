package main

import (
	"fmt"
	"log"

	"github.com/anaskhan96/soup"
)

func main() {
	resp, err := soup.Get("http://quotes.toscrape.com/tag/love/")
	if err != nil {
		log.Fatal(err)
	}
	doc := soup.HTMLParse(resp)

	var links []soup.Root
	li := doc.FindAll("div", "class", "col-md-8") // Root Node

	// Since there are two root node with same defination, the required lower Root Node must be in any one of the slice.
	// The soup library catches panics and sends log messages for the invalid node, whereas continues execution for the right one.
	for _, lin := range li {
		links = lin.FindAll("div", "class", "quote")
	}

	for _, link := range links {

		fmt.Println(link.Find("span", "class", "text").Text())
		fmt.Println("by", link.Find("small", "class", "author").Text())
		fmt.Println("----------------")
	}

}
