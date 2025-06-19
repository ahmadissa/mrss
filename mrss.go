package parser

import (
	"encoding/xml"
	"errors"
	"io"
	"net/http"
	"os"
	"strings"
)

// Root structure
type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
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

// ParseMRSS parses an MRSS XML from a local file path or HTTPS URL
func ParseMRSS(source string) (*RSS, error) {
	var reader io.ReadCloser
	var err error

	if strings.HasPrefix(source, "http://") || strings.HasPrefix(source, "https://") {
		resp, err := http.Get(source)
		if err != nil {
			return nil, err
		}
		if resp.StatusCode != 200 {
			resp.Body.Close()
			return nil, errors.New("non-200 HTTP status: " + resp.Status)
		}
		reader = resp.Body
	} else {
		reader, err = os.Open(source)
		if err != nil {
			return nil, err
		}
	}

	defer reader.Close()

	var rss RSS
	decoder := xml.NewDecoder(reader)
	err = decoder.Decode(&rss)
	return &rss, err
}
