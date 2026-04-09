package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gocolly/colly"
)

type Stock struct {
	company, price, change string
}

func main() {
	ticker := []string{
		"MSFT", "IBM", "GE", "UNP", "COST", "MCD", "V", "WMT", "DIS", "MMM",
		"INTC", "AXP", "AAPL", "BA", "CSCO", "GS", "JPM", "CRM", "VZ",
	}

	stocks := []Stock{}

	c := colly.NewCollector(
		// Set a realistic User-Agent to avoid immediate blocking
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36"),
	)

	// Add a delay to prevent Yahoo Finance from blocking with a 429 Too Many Requests
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*yahoo.com*",
		RandomDelay: 2 * time.Second,
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Printf("Something went wrong scraping %s: %s (Status: %d)\n", r.Request.URL, err, r.StatusCode)
		if r.StatusCode == 429 {
			log.Println("WARNING: Yahoo Finance is rate limiting you (429 Too Many Requests). Try again later.")
		}
	})

	// Instead of div#quote-header-info, we search the whole HTML body but restrict to the first match
	c.OnHTML("html", func(e *colly.HTMLElement) {
		stock := Stock{}

		// Use the first match to avoid concatenating multiple elements
		stock.company = e.DOM.Find("h1").First().Text()
		stock.price = e.DOM.Find("fin-streamer[data-field='regularMarketPrice']").First().Text()
		stock.change = e.DOM.Find("fin-streamer[data-field='regularMarketChangePercent']").First().Text()

		if stock.company != "" && stock.price != "" {
			fmt.Printf("Parsed -> Company: %s, Price: %s, Change: %s\n", stock.company, stock.price, stock.change)
			stocks = append(stocks, stock)
		} else {
			fmt.Printf("Warning: Could not parse stock details for %s. Layout may have changed or request was blocked.\n", e.Request.URL)
		}
	})

	for _, t := range ticker {
		c.Visit("https://finance.yahoo.com/quote/" + t + "/")
	}

	// Wait correctly waits for the visits (which are sequential here, but the limit rule applies delays)
	c.Wait()

	fmt.Println("Total Scraped Stocks:", len(stocks))

	file, err := os.Create("stocks.csv")
	if err != nil {
		log.Fatalln("Failed to create output CSV file", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"company", "price", "change"}
	writer.Write(headers)
	for _, stock := range stocks {
		record := []string{
			stock.company,
			stock.price,
			stock.change,
		}
		writer.Write(record)
	}

	fmt.Println("Finished writing data to stocks.csv")
}