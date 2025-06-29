package parser

import (
	"testing"
)

func TestParseLocalMRSSWithExpiry(t *testing.T) {
	rss, err := ParseMRSS("./mrss-with-expiry.xml")
	if err != nil {
		t.Fatalf("ParseMRSS failed: %v", err)
	}

	if rss.Channel.Title != "Sample Feed" {
		t.Errorf("Expected title 'Sample Feed', got '%s'", rss.Channel.Title)
	}

	if len(rss.Channel.Items) != 4 {
		t.Errorf("Expected 4 items, got %d", len(rss.Channel.Items))
	}

	for i, item := range rss.Channel.Items {
		if item.Valid == nil {
			t.Errorf("Expected dcterms:valid in item %d", i+1)
		}
		for _, media := range item.MediaContents {
			if media.URL == "" {
				t.Errorf("Missing media:content URL in item %d", i+1)
			}
			if media.Valid == nil {
				t.Errorf("Expected dcterms:valid in media:content of item %d", i+1)
			}
		}
	}
}

func TestParseLocalMRSSWithoutExpiry(t *testing.T) {
	rss, err := ParseMRSS("./mrss-feed-no-expiry.xml")
	if err != nil {
		t.Fatalf("ParseMRSS failed: %v", err)
	}

	if len(rss.Channel.Items) != 4 {
		t.Errorf("Expected 4 items, got %d", len(rss.Channel.Items))
	}

	for i, item := range rss.Channel.Items {
		if item.Valid != nil {
			t.Errorf("Expected no dcterms:valid in item %d", i+1)
		}
		for _, media := range item.MediaContents {
			if media.URL == "" {
				t.Errorf("Missing media:content URL in item %d", i+1)
			}
			if media.Valid != nil {
				t.Errorf("Expected no dcterms:valid in media:content of item %d", i+1)
			}
		}
	}
}

func TestParseRemoteMRSSWithExpiry(t *testing.T) {
	url := "https://files.cloud-digitalsignage.com/mrss/mrss-with-expiry.xml"
	rss, err := ParseMRSS(url)
	if err != nil {
		t.Fatalf("Failed to fetch remote MRSS with expiry: %v", err)
	}

	if len(rss.Channel.Items) == 0 {
		t.Errorf("Expected non-zero items in remote feed")
	}

	for i, item := range rss.Channel.Items {
		if item.Valid == nil {
			t.Errorf("Expected dcterms:valid in item %d", i+1)
		}
		for _, media := range item.MediaContents {
			if media.URL == "" {
				t.Errorf("Missing media URL in item %d", i+1)
			}
			if media.Valid == nil {
				t.Errorf("Expected dcterms:valid in media content %d", i+1)
			}
		}
	}
}

func TestParseRemoteMRSSWithoutExpiry(t *testing.T) {
	url := "https://files.cloud-digitalsignage.com/mrss/mrss-feed-no-expiry.xml"
	rss, err := ParseMRSS(url)
	if err != nil {
		t.Fatalf("Failed to fetch remote MRSS without expiry: %v", err)
	}

	if len(rss.Channel.Items) == 0 {
		t.Errorf("Expected non-zero items in remote feed")
	}

	for i, item := range rss.Channel.Items {
		if item.Valid != nil {
			t.Errorf("Expected no dcterms:valid in item %d", i+1)
		}
		for _, media := range item.MediaContents {
			if media.URL == "" {
				t.Errorf("Missing media URL in item %d", i+1)
			}
			if media.Valid != nil {
				t.Errorf("Expected no dcterms:valid in media content %d", i+1)
			}
		}
	}
}

func TestParseLocalJSONWithExpiry(t *testing.T) {
	rss, err := ParseJSONFeed("./mrss-with-expiry.json")
	if err != nil {
		t.Fatalf("ParseJSONFeed failed: %v", err)
	}

	if rss.Channel.Title != "Sample Feed" {
		t.Errorf("Expected title 'Sample Feed', got '%s'", rss.Channel.Title)
	}

	if len(rss.Channel.Items) != 4 {
		t.Errorf("Expected 4 items, got %d", len(rss.Channel.Items))
	}

	for i, item := range rss.Channel.Items {
		if item.Valid == nil {
			t.Errorf("Expected valid in item %d", i+1)
		}
		for _, media := range item.MediaContents {
			if media.URL == "" {
				t.Errorf("Missing media URL in item %d", i+1)
			}
			if media.Valid == nil {
				t.Errorf("Expected valid in media content %d", i+1)
			}
		}
	}
}

func TestParseLocalJSONWithoutExpiry(t *testing.T) {
	rss, err := ParseJSONFeed("./mrss-feed-no-expiry.json")
	if err != nil {
		t.Fatalf("ParseJSONFeed failed: %v", err)
	}

	if len(rss.Channel.Items) != 4 {
		t.Errorf("Expected 4 items, got %d", len(rss.Channel.Items))
	}

	for i, item := range rss.Channel.Items {
		if item.Valid != nil {
			t.Errorf("Expected no valid in item %d", i+1)
		}
		for _, media := range item.MediaContents {
			if media.URL == "" {
				t.Errorf("Missing media URL in item %d", i+1)
			}
			if media.Valid != nil {
				t.Errorf("Expected no valid in media content %d", i+1)
			}
		}
	}
}
