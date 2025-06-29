package parser

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"net/http"
	"os"
	"strings"
)

// Root structure
type RSS struct {
	XMLName xml.Name `xml:"rss" json:"-"`
	Channel Channel  `xml:"channel" json:"channel"`
}

type Channel struct {
	Title       string `xml:"title" json:"title"`
	Description string `xml:"description" json:"description"`
	Items       []Item `xml:"item" json:"items"`
}

type Item struct {
	Title         string         `xml:"title" json:"title"`
	Description   string         `xml:"description" json:"description"`
	Valid         *string        `xml:"http://purl.org/dc/terms/ valid,omitempty" json:"valid,omitempty"`
	MediaContents []MediaContent `xml:"http://search.yahoo.com/mrss/ content" json:"mediaContents"`
}

type MediaContent struct {
	URL       string  `xml:"url,attr" json:"url"`
	Type      string  `xml:"type,attr" json:"type"`
	Medium    string  `xml:"medium,attr" json:"medium"`
	Duration  string  `xml:"duration,attr" json:"duration"`
	ChangeKey string  `xml:"change_key,attr,omitempty" json:"changeKey,omitempty"`
	Valid     *string `xml:"http://purl.org/dc/terms/ valid,omitempty" json:"valid,omitempty"`
}

func (m MediaContent) GetChangeKey() string {
	input := m.URL
	if m.ChangeKey != "" {
		input += m.ChangeKey
	}
	hash := md5.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
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

// ParseJSONFeed parses a JSON-based feed from a local file path or HTTPS URL
func ParseJSONFeed(source string) (*RSS, error) {
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
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&rss)
	return &rss, err
}
