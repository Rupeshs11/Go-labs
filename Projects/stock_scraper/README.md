# 📈 Stock Scraper

An aggressive, highly-concurrent web scraping tool built in **Go** that harvests real-time stock prices directly from **Yahoo Finance**. The scraped data is sanitized and exported cleanly into a CSV format.

---

## 🛠 Features

- **Blazing Fast Scraping**: Driven by the lightning-fast `gocolly/colly` scraping framework.
- **Anti-Bot Protection Bypass**: Bypasses basic 429 Too Many Requests blocking mechanisms:
  - Spoofs authentic User-Agents.
  - Implements an asynchronous `LimitRule`, forcing organic 2-second delays between individual HTTP requests to Yahoo servers.
- **Automated CSV Generation**: Dumps the harvested data seamlessly into `stocks.csv` using Go's built-in `encoding/csv` standard library.

---

## 💻 Code Explanations

### Setting up the Collector (`colly`)

```go
c := colly.NewCollector(
    colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64)..."),
)

c.Limit(&colly.LimitRule{
    DomainGlob:  "*yahoo.com*",
    RandomDelay: 2 * time.Second,
})
```

- Instantiates a generic scraper configuration setting the framework globally.
- The `colly.LimitRule` acts as a crucial circuit breaker, introducing hard-coded delays to ensure Yahoo servers don't actively blacklist the executing IP.

### Scraping Logic (DOM Traversal)

```go
c.OnHTML("html", func(e *colly.HTMLElement) {
    stock := Stock{}

    // Locate and extract deeply nested structural nodes
    stock.company = e.DOM.Find("h1").First().Text()
    stock.price = e.DOM.Find("fin-streamer[data-field='regularMarketPrice']").First().Text()
    stock.change = e.DOM.Find("fin-streamer[data-field='regularMarketChangePercent']").First().Text()
    // ...
})
```

- Colly hooks directly onto the `html` root payload.
- GoQuery-style DOM Selectors (`e.DOM.Find`) traverse the HTML nodes to pinpoint the specific `fin-streamer` tags holding the absolute current market price and subsequent yield changes.

### CSV Data Flushing

```go
file, _ := os.Create("stocks.csv")
writer := csv.NewWriter(file)
defer writer.Flush()

writer.Write([]string{"company", "price", "change"})
// Logic wrapping records and pushing them...
```

- Employs native `os` features to create files and wraps them into `csv.NewWriter()`, building the header string arrays and cascading data inside.

---

## 🚀 How to Run

1. Clone the repository and navigate to the project directory:

   ```bash
   cd Projects/stock_scraper
   ```

2. Download dependencies (`github.com/gocolly/colly`):

   ```bash
   go mod tidy
   ```

3. Run the application (This may take roughly `~30-40 seconds` as the delay rules throttle parsing to be safe):

   ```bash
   go run main.go
   ```

4. A new file `stocks.csv` will be generated in your immediate directory filled with Wall Street metrics!

### Optional Test Request

A secondary file `fetch.go` is included to raw-dump a single payload page.

```bash
go run fetch.go
# Outputs a massive payload stream to msft.html
```

---

## Dependencies

| Package                                             | Purpose                                                                       |
| --------------------------------------------------- | ----------------------------------------------------------------------------- |
| [`gocolly/colly`](https://github.com/gocolly/colly) | Bleeding edge web scraping framework targeting speed and asynchronous scaling |
