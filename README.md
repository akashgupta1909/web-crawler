# 🕸️ Web Crawler (Golang)

A high-performance internal web crawler written in **Go**, built to recursively crawl links on a website with customizable concurrency and page limits.

---

## 🚀 Features

- 🕵️ Crawls only internal pages (same domain)
- 🌐 Resolves both relative and absolute URLs
- 📈 Crawl summary report (pages sorted by frequency)
- ✅ Well-structured tests for core modules

---

## 📦 Project Structure

```plaintext
.
├── main.go                      # CLI entry point
├── configure.go                # Configuration and setup
├── crawlPage.go                # Crawling logic
├── internal
│   ├── customHTML              # HTML parsing and link extraction
│   │   ├── getHTML.go
│   │   ├── getURLsFromHTML.go
│   │   └── getURLsFromHTML_test.go
│   ├── customPrint             # Report formatting
│   │   ├── printReport.go
│   │   └── printReport_test.go
│   └── customURL               # URL normalization
│       ├── normalize_url.go
│       └── normalize_url_test.go
├── go.mod
└── go.sum
```

## 🛠️ Installation

```plaintext
git clone git@github.com:akashgupta1909/web-crawler.git
cd web-crawler
go mod tidy
```

## 🧪 Testing

Run the tests using the following command:

```plaintext
go test ./...

```

## 🏗️ Build

To build the project, use the following command:

```plaintext
go build
```

## 🏃‍♂️ Usage

```plaintext
go run main.go <base-url> [maxConcurrency] [maxPages]
```

### Example

```plaintext
go run main.go https://example.com 10 100
```

This command will crawl the website `https://example.com` with a maximum of 10 concurrent requests and a limit of 100 unique urls.

### Sample Output

```plaintext
starting crawl of: https://example.blog.dev
concurrency 10
max pages 20
base url https://example.blog.dev
starting crawaling on: blog.boot.dev/path
...

==========================
Report for: https://example.blog.dev
==========================
Found 5 internal links to example.blog.dev
Found 3 internal links to example.blog.dev/about
Found 2 internal links to example.blog.dev/path
...
```
