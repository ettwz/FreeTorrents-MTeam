package main

import (
	"encoding/json"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestStringNumberUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		raw  string
		want string
	}{
		{name: "string", raw: `{"category":"442"}`, want: "442"},
		{name: "number", raw: `{"category":442}`, want: "442"},
		{name: "null", raw: `{"category":null}`, want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var torrent TorrentInfo
			if err := json.Unmarshal([]byte(tt.raw), &torrent); err != nil {
				t.Fatalf("json.Unmarshal() error = %v", err)
			}
			if got := string(torrent.Category); got != tt.want {
				t.Fatalf("category = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestCategoryBlocked(t *testing.T) {
	blockList := []StringNumber{"442", " 444 "}

	if !categoryBlocked("442", blockList) {
		t.Fatal("expected category 442 to be blocked")
	}
	if !categoryBlocked("444", blockList) {
		t.Fatal("expected trimmed category 444 to be blocked")
	}
	if categoryBlocked("44", blockList) {
		t.Fatal("expected partial category 44 to pass")
	}
	if categoryBlocked("", blockList) {
		t.Fatal("expected empty category to pass")
	}
}

func TestSizeBlocked(t *testing.T) {
	tests := []struct {
		name    string
		size    float64
		minSize float64
		maxSize float64
		want    bool
	}{
		{name: "below min", size: 1.9, minSize: 2, maxSize: 10, want: true},
		{name: "at min", size: 2, minSize: 2, maxSize: 10, want: false},
		{name: "within range", size: 5, minSize: 2, maxSize: 10, want: false},
		{name: "at max", size: 10, minSize: 2, maxSize: 10, want: false},
		{name: "above max", size: 10.1, minSize: 2, maxSize: 10, want: true},
		{name: "disabled min", size: 0.1, minSize: 0, maxSize: 10, want: false},
		{name: "disabled max", size: 100, minSize: 2, maxSize: 0, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sizeBlocked(tt.size, tt.minSize, tt.maxSize); got != tt.want {
				t.Fatalf("sizeBlocked(%v, %v, %v) = %v, want %v", tt.size, tt.minSize, tt.maxSize, got, tt.want)
			}
		})
	}
}

func TestConfiguredPageSize(t *testing.T) {
	tests := []struct {
		name     string
		pageSize int
		want     int
	}{
		{name: "configured", pageSize: 50, want: 50},
		{name: "zero uses default", pageSize: 0, want: defaultPageSize},
		{name: "negative uses default", pageSize: -1, want: defaultPageSize},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := configuredPageSize(tt.pageSize); got != tt.want {
				t.Fatalf("configuredPageSize(%d) = %d, want %d", tt.pageSize, got, tt.want)
			}
		})
	}
}

func TestConfUnmarshalCategoryBlockList(t *testing.T) {
	raw := []byte("categoryBlockList:\n  - 442\n  - \"444\"\n")

	var conf Conf
	if err := yaml.Unmarshal(raw, &conf); err != nil {
		t.Fatalf("yaml.Unmarshal() error = %v", err)
	}

	if got := string(conf.CategoryBlockList[0]); got != "442" {
		t.Fatalf("categoryBlockList[0] = %q, want %q", got, "442")
	}
	if got := string(conf.CategoryBlockList[1]); got != "444" {
		t.Fatalf("categoryBlockList[1] = %q, want %q", got, "444")
	}
}
