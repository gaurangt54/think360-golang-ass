package main

import (
	"fmt"
	"thinktask/tasks"

	"github.com/gocolly/colly"
)

func main() {

	var option int
	fmt.Printf("Enter 1: FizzBuzz, 2: ATM, 3: Web Scrap -> ")
	fmt.Scanln(&option)
	switch option {
	case 1:
		tasks.Task1()

	case 2:
		tasks.Task2()

	case 3:

		var keyword string
		fmt.Printf("Enter your keyword -> ")
		fmt.Scanln(&keyword)

		c := colly.NewCollector(colly.AllowedDomains("www.amazon.in"))

		c.OnRequest(func(r *colly.Request) {
			fmt.Println("Link of the page:", r.URL)
		})

		c.OnHTML("div.s-result-list.s-search-results.sg-row", func(h *colly.HTMLElement) {
			h.ForEach("div.a-section.a-spacing-base", func(_ int, h *colly.HTMLElement) {
				var name string
				name = h.ChildText("span.a-size-base-plus.a-color-base.a-text-normal")
				// var stars string
				// stars = h.ChildText("span.a-icon-alt")
				// var price string
				// price = h.ChildText("span.a-price-whole")

				fmt.Println(name)
				fmt.Println()

			})
		})

		var link = "https://www.amazon.in/s?k=" + keyword

		c.Visit(link)

	default:
		fmt.Print("Enter Correct Option")
	}

}
