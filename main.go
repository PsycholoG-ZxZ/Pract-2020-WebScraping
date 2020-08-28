package main

import (
	"log"

	"github.com/gocolly/colly"
)

var (
	auth_key = "0"
)

func main() {
	// create a new collector
	c := colly.NewCollector()

	//authenticate
	err := c.Post("http://quotes.toscrape.com/login", map[string]string{"username": "CheckCheck", "password": "qweasd123"})
	if err != nil {
		log.Fatal(err)
	}

	// attach callbacks after login
	c.OnResponse(func(r *colly.Response) {
		log.Println("response received", r.StatusCode)
	})

	// start scraping
	frst_elem := 0

	c.OnHTML("div.col-md-4", func(e *colly.HTMLElement) {
		key := e.Text
		frst_elem = frst_elem + 1
		if frst_elem == 1 {
			log.Println("check: " + key)
		}

	})

	c.Visit("http://quotes.toscrape.com/")

}
