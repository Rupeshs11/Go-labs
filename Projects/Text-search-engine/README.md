# 🔎 Text Search Engine

A high-performance full-text search engine built in **Go**. This engine efficiently ingests compressed Wikipedia XML dumps (`.xml.gz`), constructs a lightning-fast **Inverted Index**, and allows instant querying utilizing advanced text-filtering and stemming!

---

## 🛠 Features

- **Inverted Indexing**: Rapidly builds a dictionary linking individual tokens (words) to their respective Document IDs to achieve `O(1)` query lookups.
- **Set Intersection Queries**: Employs an ultra-fast algorithm to cross-reference multiple IDs when searching for complex queries with multiple words.
- **Advanced Text Normalization**:
  - **Snowball Stemming**: Incorporates the Snowball algorithm (`kljensen/snowball`) to strip language suffixes (e.g., matching "fishing" -> "fish").
  - **Stopword Filtering**: Automatically ignores common filler words like `the`, `and`, `to` via `O(1)` map validation.
  - **Lowercasing**: Unifies data structures by squashing text uniformly to prevent casing misses.

---

## 💻 Code Explanations

### Building the Inverted Index

```go
type Index map[string][]int

func (idx Index) Add(docs []document) {
    for _, doc := range docs {
        for _, token := range analyze(doc.Text) {
            // Append Document ID to Token's slice securely
            idx[token] = append(idx[token], doc.ID)
        }
    }
}
```

- Core indexing system based on a simple mapped data structure matching text strings explicitly to integer slices containing internal Document IDs.

### Search Engine Intersections (`index.go`)

```go
func (idx Index) Search(text string) []int {
    var r []int
    for _, token := range analyze(text) {
        if ids, ok := idx[token]; ok {
            if r == nil {
                r = ids
            } else {
                r = Intersection(r, ids) // Slices intersecting IDs
            }
        } else {
            return nil
        }
    }
    return r
}
```

- Iterates over each token within your search parameter.
- Employs an intelligent intersection algorithm tracking two arrays, validating if subsequent keywords appear in the _exact same underlying documents_, drastically cutting down false positives!

### Text Filter Pipelines (`filter.go`)

```go
func stemmerFilter(tokens []string) []string {
    r := make([]string, len(tokens))
    for i, token := range tokens {
        // Run tokens through the Snowball english stemmer
        r[i] = snowballeng.Stem(token, false)
    }
    return r
}
```

- Integrates cleanly with github modules to physically truncate english words down to their root meaning (e.g. `cats` -> `cat`).

---

## 🚀 How to Run

1. Clone the repository and navigate to the project directory:

   ```bash
   cd Projects/Text-search-engine
   ```

2. Tidy dependencies (`github.com/kljensen/snowball`):

   ```bash
   go mod tidy
   ```

3. Extract & Search! By default, the application is set up to parse `enwiki-latest-abstract1.xml.gz` against the query `Small wild cat`:
   ```bash
   go run main.go
   # Alternatively query specifically by flags:
   # go run main.go -p data.xml.gz -q "query strings"
   ```

---

## Dependencies

| Package                                                     | Purpose                                                                                |
| ----------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| [`kljensen/snowball`](https://github.com/kljensen/snowball) | Industry-standard stemmer reducing English words back to their root suffix/prefix core |
