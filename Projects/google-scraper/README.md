# 🔍 Google Scraper

A powerful, configurable Google Search Results scraper built in **Go** using the `goquery` HTML parsing library. It supports **190+ country-specific Google domains**, randomized User-Agent rotation, optional proxy routing, and paginated result extraction to return ranked search results.

---

## 🛠 Features

- **190+ Google Domains**: Supports country-code-specific Google domains (e.g., `google.co.in`, `google.de`, `google.co.jp`, etc.) for localized results.
- **Randomized User-Agents**: Rotates across 6 different browser User-Agent strings on every request to avoid fingerprinting detection.
- **Proxy Support**: Optional proxy routing via `http.Transport` for anonymized scraping.
- **Pagination**: Automatically constructs multiple page URLs using `start` and `num` query parameters to scrape beyond the first results page.
- **Backoff Delays**: Configurable sleep intervals between requests to prevent IP bans from Google.
- **Structured Output**: Returns clean `SearchResult` structs containing rank, URL, title, and description.

---

## 💻 Code Explanations

### Building Google URLs (`buildGoogleUrls`)

```go
func buildGoogleUrls(searchTerm, countryCode, languageCode string, pages, count int) ([]string, error) {
    searchTerm = strings.Replace(searchTerm, " ", "+", -1)
    if googleBase, found := googleDomains[countryCode]; found {
        for i := 0; i < pages; i++ {
            start := i * count
            scrapeURL := fmt.Sprintf("%s%s&num=%d&hl=%s&start=%d&filter=0",
                          googleBase, searchTerm, count, languageCode, start)
            toScrape = append(toScrape, scrapeURL)
        }
    }
    return toScrape, nil
}
```

- Looks up the country-specific Google base URL from the massive `googleDomains` map.
- Constructs paginated URLs using the `start` offset and `num` count to fetch custom result batches.

### HTML Parsing (`googleResultParsing`)

```go
doc, _ := goquery.NewDocumentFromResponse(response)
sel := doc.Find("div.g")

for i := range sel.Nodes {
    item := sel.Eq(i)
    link, _ := item.Find("a").Attr("href")
    title := item.Find("h3").Text()
    desc := item.Find(".VwiC3b, .st").Text()
    // ...
}
```

- Parses the raw HTTP response using `goquery.NewDocumentFromResponse`.
- Targets the classic `div.g` container used by Google for individual search result entries.
- Extracts the link (`a[href]`), title (`h3`), and description (`.VwiC3b` or `.st`) from each result node.

### Proxy-Aware HTTP Client (`getScrapeClient`)

```go
func getScrapeClient(proxyString interface{}) *http.Client {
    switch v := proxyString.(type) {
    case string:
        proxyUrl, _ := url.Parse(v)
        return &http.Client{
            Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)},
            Timeout:   10 * time.Second,
        }
    default:
        return &http.Client{Timeout: 10 * time.Second}
    }
}
```

- Uses a Go **type switch** (`interface{}`) to dynamically determine whether to route traffic through a proxy or use a direct connection.
- Demonstrates a real-world example of Go's flexible `interface{}` pattern for optional parameters.

---

## 🚀 How to Run

1. Clone the repository and navigate to the project directory:

   ```bash
   cd Projects/google-scraper
   ```

2. Download dependencies (`goquery`):

   ```bash
   go mod tidy
   ```

3. Run the scraper:
   ```bash
   go run main.go
   ```

### Customizing Search Parameters

Edit the `main()` function to change the search query, country, language, and pagination:

```go
// GoogleScrape(searchTerm, countryCode, languageCode, proxyString, pages, count, backoff)
res, err := GoogleScrape("golang tutorials", "in", "en", nil, 2, 10, 10)
```

| Parameter      | Description                                     |
| -------------- | ----------------------------------------------- |
| `searchTerm`   | The search query string                         |
| `countryCode`  | Google domain country code (e.g., `com`, `in`)  |
| `languageCode` | Language for results (e.g., `en`, `de`)         |
| `proxyString`  | Proxy URL string or `nil` for direct connection |
| `pages`        | Number of result pages to scrape                |
| `count`        | Number of results per page                      |
| `backoff`      | Seconds to wait between page requests           |

---

## Dependencies

| Package                                             | Purpose                                                |
| --------------------------------------------------- | ------------------------------------------------------ |
| [`goquery`](https://github.com/PuerkitoBio/goquery) | jQuery-like HTML DOM traversal and manipulation for Go |
