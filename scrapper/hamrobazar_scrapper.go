package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/yhat/scrape"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func main() {
	// category id; differs according to domains
	var catID int
	noOfResults := flag.Int("no_of_results", 20, "Please specify in order of 20")
	domain := flag.String("domain", "laptop", "Specify between domains: Laptop, Business or Mobile")
	sortBy := flag.String("sort_by", "popular", "Specify between popular or latest")
	flag.Parse()

	if *noOfResults%20 != 0 {
		fmt.Fprintf(os.Stderr, "no_of_results: %d not divisible by 20\n", *noOfResults)
		flag.Usage()
		os.Exit(1)
	}
	*domain = strings.ToLower(*domain)
	domainList := map[string]bool{
		"laptop":   true,
		"business": true,
		"mobile":   true,
	}

	if _, ok := domainList[*domain]; !ok {
		fmt.Fprintf(os.Stderr, "Incorrect domain: %s\n", *domain)
		flag.Usage()
		os.Exit(1)
	}

	switch *domain {
	case "laptop":
		*domain = "c22-computer-and-peripherals-laptops"
		catID = 22
	case "mobile":
		*domain = "c31-mobile-and-accessories-handsets"
		catID = 31
	case "business":
		*domain = "c191-business-and-industrial-business-for-sale"
		catID = 191
	}

	*sortBy = strings.ToLower(*sortBy)
	sortList := map[string]bool{
		"popular": true,
		"latest":  true,
	}
	if _, ok := sortList[*sortBy]; !ok {
		fmt.Fprintf(os.Stderr, "Incorrect sort: %s\n", *sortBy)
		flag.Usage()
		os.Exit(1)
	}

	switch *sortBy {
	case "popular":
		*sortBy = "popularad"
	case "latest":
		*sortBy = "siteid"
	}

	// Creating a csv file
	file, err := os.Create("file.csv")
	if err != nil {
		log.Fatal("Could not create file")
	}
	defer file.Close()

	// Writing headers to it
	headers := []string{"title", "specs", "description", "name", "address", "date", "price"}
	writer := csv.NewWriter(file)
	writer.Write(headers)

	baseURL := "http://hamrobazaar.com/%s?catid=%d&order=%s&offset=%d"
	var records [][]string
	for i := 0; i < *noOfResults; i = i + 20 {
		log.Printf("Fetching records from %d to %d", i+1, i+20)
		url := fmt.Sprintf(baseURL, *domain, catID, *sortBy, i)
		records = append(records, Fetch(url)...)

	}

	writer.WriteAll(records)
	writer.Flush()
	if err := writer.Error(); err != nil {
		log.Fatal(err)
	}
}

func Fetch(url string) [][]string {
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	root, err := html.Parse(resp.Body)

	if err != nil {
		log.Fatal("could not parse")
	}

	matcher := func(n *html.Node) bool {
		if n.FirstChild != nil && n.Parent != nil && n.Parent.Parent != nil && n.Parent.Parent.Parent != nil {
			if scrape.Attr(n, "height") == "110" && n.FirstChild.DataAtom == atom.A && n.Parent.Parent.Parent.DataAtom == atom.Table {
				return true
			}
		}
		return false
	}
	records := [][]string{}
	articles := scrape.FindAll(root, matcher)
	for i, article := range articles {
		record := make([]string, 7)

		// firstchild node
		n1 := article.FirstChild
		// title and specs
		record[0] = scrape.Text(n1)
		record[1] = scrape.Text(n1.NextSibling.NextSibling)

		// description node
		n1 = n1.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling
		record[2] = n1.Data
		if i < 2 {
			record[2] = n1.PrevSibling.Data
		}

		// name and address
		record[3] = n1.NextSibling.NextSibling.NextSibling.NextSibling.FirstChild.Data
		record[4] = article.LastChild.FirstChild.Data

		// date and price
		record[5] = article.NextSibling.NextSibling.FirstChild.Data
		record[6] = scrape.Text(article.NextSibling.NextSibling.NextSibling.NextSibling)
		records = append(records, record)
	}
	return records
}
