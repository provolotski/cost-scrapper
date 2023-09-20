package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

type products struct {
	url, name, price string
}

func main() {

	c := colly.NewCollector(colly.AllowedDomains("www.21vek.by"))

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting: ", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Page visited: ", r.Request.URL)
	})
	c.OnHTML("li.result__item", func(e *colly.HTMLElement) {

		prod := products{}
		prod.url = e.ChildAttr("a", "href")
		prod.name = e.ChildText(".result__name")

		prod.price = e.ChildText(".g-item-data")
		//prod.price = e.ChildAttr(".result__price", "content")
		//prod.price = e.ChildText("content")

		//fmt.Println(prod.url)
		//fmt.Println(e.ChildText(".result__price"))
		fmt.Println(prod.name)
		fmt.Println(prod.price)
	})
	c.OnScraped(func(r *colly.Response) {
		fmt.Println(r.Request.URL, " scrapped!")
	})
	c.Visit("https://www.21vek.by/refrigerators/")
}

//li.result__item:nth-child(13) > dl:nth-child(2) > div:nth-child(1) > dt:nth-child(1) > a:nth-child(2) > span:nth-child(1) > div:nth-child(1)
