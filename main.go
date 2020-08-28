package main

import (
	"flag"
	"log"
	"time"

	"github.com/gocolly/colly"
)

var (
	auth_key = "0"
)

func authenticate(c *colly.Collector) {
	err := c.Post("http://quotes.toscrape.com/login", map[string]string{"username": "CheckCheck", "password": "qweasd123"})
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// create a new collector

	//flags
	boolPtr := flag.Bool("online", false, "a bool")
	timePtr := flag.Float64("MIN", 1, "a float64")

	flag.Parse()

	i := 0

	log.Println(*boolPtr)
	log.Println(*timePtr)
	//*timePtr = (*timePtr) * 0.1 // DELETE AFTER DEBUGING

	for i < 1 {
		c := colly.NewCollector()

		authenticate(c)

		// attach callbacks after login
		c.OnResponse(func(r *colly.Response) {
			log.Println("response received", r.StatusCode)
		})

		// start scraping
		frstElem := 0

		c.OnHTML("div.col-md-4", func(e *colly.HTMLElement) {
			key := e.Text
			frstElem = frstElem + 1
			if frstElem == 1 {
				log.Println("check: " + key)

			}

		})

		c.Visit("http://quotes.toscrape.com/")

		if *boolPtr == false {
			break
		}
		time.Sleep(time.Duration(*timePtr) * time.Minute)

	}

}
