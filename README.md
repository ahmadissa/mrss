# MRSS Parser for Go

This is a lightweight Go module to parse **Media RSS (MRSS)** feeds with optional `dcterms:valid` fields and `media:content` elements. It supports both **local files and HTTPS URLs**, and is built using Go’s standard `encoding/xml` package for speed, simplicity, and full control over namespaced elements.

## 📦 Features

- Parses `<media:content>` elements (with attributes like `url`, `type`, `duration`, etc.)
- Handles optional `<dcterms:valid>` fields at both item and media level
- Accepts both local file paths and HTTPS URLs
- Fully compatible with feeds containing images, videos, or mixed media
- Returns Go structs ready for use

## 📂 Example Feed Supported

### With Expiry
```xml
<item>
  <title>Item 1</title>
  <description>Media RSS Item</description>
  <dcterms:valid>start=2023-01-01T00:00:00;end=2024-01-01T00:00:00;scheme=W3C-DTF</dcterms:valid>
  <media:content url="http://server/image1.jpg" type="image/jpeg" medium="image" duration="10">
    <dcterms:valid>start=2023-01-01T00:00:00;end=2024-01-01T00:00:00;scheme=W3C-DTF</dcterms:valid>
  </media:content>
</item>
```

### Without Expiry
```xml
<item>
  <title>Item 2</title>
  <description>Media RSS Item</description>
  <media:content url="http://server/image2.jpg" type="image/jpeg" medium="image" duration="10" />
</item>
```

---

## 🧑‍💻 Usage

### Step 1: Import the parser

```go
import "github.com/yourusername/mrssparser/parser"
```

### Step 2: Parse a Local File or HTTPS URL

```go
rss, err := parser.ParseMRSS("local-file.xml")
// or
rss, err = parser.ParseMRSS("https://example.com/feed.xml")

if err != nil {
  log.Fatal(err)
}

for _, item := range rss.Channel.Items {
  fmt.Println(item.Title, item.Description)
  for _, media := range item.MediaContents {
    fmt.Println("  URL:", media.URL, "Type:", media.Type)
  }
}
```

---

## 📁 Struct Overview

```go
type RSS struct {
  Channel Channel `xml:"channel"`
}

type Channel struct {
  Title       string `xml:"title"`
  Description string `xml:"description"`
  Items       []Item `xml:"item"`
}

type Item struct {
  Title         string         `xml:"title"`
  Description   string         `xml:"description"`
  Valid         *string        `xml:"http://purl.org/dc/terms/ valid,omitempty"`
  MediaContents []MediaContent `xml:"http://search.yahoo.com/mrss/ content"`
}

type MediaContent struct {
  URL      string  `xml:"url,attr"`
  Type     string  `xml:"type,attr"`
  Medium   string  `xml:"medium,attr"`
  Duration string  `xml:"duration,attr"`
  Valid    *string `xml:"http://purl.org/dc/terms/ valid,omitempty"`
}
```

---

## 🧪 Testing

You can run the built-in tests for local and remote MRSS parsing:

```bash
go test ./parser
```

This includes:
- ✅ `mrss-feed-no-expiry.xml`
- ✅ `mrss-with-expiry.xml`
- ✅ Remote feeds like:
  - `https://files.cloud-digitalsignage.com/mrss/mrss-feed-no-expiry.xml`
  - `https://files.cloud-digitalsignage.com/mrss/mrss-with-expiry.xml`

---

## 🔖 License

MIT — feel free to use and modify as needed.

---
