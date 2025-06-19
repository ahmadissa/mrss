# MRSS Parser for Go

This is a lightweight Go module to parse **Media RSS (MRSS)** feeds with optional `dcterms:valid` fields and `media:content` elements. It's built using Go’s standard `encoding/xml` package for speed, simplicity, and full control over namespaced elements.

## 📦 Features

- Parses `<media:content>` elements (with attributes like `url`, `type`, `duration`, etc.)
- Handles optional `<dcterms:valid>` fields at both item and media level
- Works with both MRSS feeds that include expiry and those that don’t
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

### Step 2: Parse a Feed

```go
f, _ := os.Open("feed.xml")
rss, err := parser.ParseMRSS(f)
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

To test with your own files:

```bash
go run main.go simple-media-rss-with-expiry.xml
```

Or adapt the feed paths as needed.

---

## 🔖 License

MIT — feel free to use and modify as needed.

---

## ✍️ Author

Ahmad Issa – [EasySignage](https://easysignage.com)
