package main

import (
	"fmt"
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"time"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"gopkg.in/mgo.v2"
)

const (
	// DBAddress ...
	DBAddress = "localhost"
)

// News ...
type News struct {
	ID          bson.ObjectId `json:"_id" bson:"_id"`
	Title       string        `json:"title" bson:"title"`
	Description string        `json:"description" bson:"description"`
	By          string        `json:"by" bson:"by"`
	URL         string        `json:"url" bson:"url"`
}

// NewsArray ...
type NewsArray []News

func main() {

	session, err := mgo.Dial(DBAddress)
	if err != nil {
		log.Fatal("Error connecting database", err)
	}
	defer session.Close()
	col := session.DB("News").C("onlinekhabar")
	for {

		resp, err := http.Get("http://www.onlinekhabar.com/content/news/")

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
			if n.DataAtom == atom.A && n.Parent != nil && n.Parent.Parent != nil {
				if n.Parent.DataAtom == atom.H2 {
					return scrape.Attr(n.Parent.Parent, "class") == "news_loop"
				}
			}
			return false
		}

		articles := scrape.FindAll(root, matcher)
		for _, article := range articles {

			title := scrape.Text(article)
			description := scrape.Text(article.Parent.NextSibling)
			id := bson.NewObjectId()
			url := fmt.Sprint("localhost:8080", "/dbs/", id.Hex())
			by := "http://www.onlinekhabar.com"
			news := News{
				ID:          id,
				Title:       title,
				Description: description,
				By:          by,
				URL:         url,
			}
			err := col.Insert(news)
			if err != nil {
				log.Fatal("Error inserting into database", err)
			}

		}
		time.Sleep(10 * time.Minute)
	}
}
