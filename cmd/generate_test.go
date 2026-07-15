package cmd

import (
	"strings"
	"testing"
)

func TestGenerateName(t *testing.T) {
	adjectives := []string{"quick"}
	nouns := []string{"fox"}

	tests := []struct {
		name         string
		prefix       string
		separator    string
		appendRandom bool
		randomChars  string
		randomLength int
		want         string
	}{
		{name: "basic", separator: "-", want: "quick-fox"},
		{name: "custom separator", separator: "_", want: "quick_fox"},
		{name: "with prefix", prefix: "app", separator: "-", want: "app-quick-fox"},
		{name: "random suffix disabled without chars", separator: "-", appendRandom: true, randomChars: "", randomLength: 3, want: "quick-fox"},
		{name: "random suffix disabled with zero length", separator: "-", appendRandom: true, randomChars: "0123456789", randomLength: 0, want: "quick-fox"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateName(adjectives, nouns, tt.prefix, tt.separator, tt.appendRandom, tt.randomChars, tt.randomLength)
			if got != tt.want {
				t.Errorf("generateName() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestGenerateNameRandomSuffix(t *testing.T) {
	got := generateName([]string{"quick"}, []string{"fox"}, "", "-", true, "x", 4)
	want := "quick-fox-xxxx"
	if got != want {
		t.Errorf("generateName() = %q, want %q", got, want)
	}

	got = generateName([]string{"quick"}, []string{"fox"}, "", "-", true, "0123456789", 3)
	parts := strings.Split(got, "-")
	if len(parts) != 3 {
		t.Fatalf("expected 3 parts, got %d in %q", len(parts), got)
	}
	if len(parts[2]) != 3 {
		t.Errorf("expected suffix of length 3, got %q", parts[2])
	}
	for _, c := range parts[2] {
		if c < '0' || c > '9' {
			t.Errorf("expected numeric suffix, got %q", parts[2])
		}
	}
}

func TestGenerateNames(t *testing.T) {
	names := generateNames([]string{"quick"}, []string{"fox"}, "", "-", false, "", 0, 5)
	if len(names) != 5 {
		t.Fatalf("expected 5 names, got %d", len(names))
	}
	for _, name := range names {
		if name != "quick-fox" {
			t.Errorf("expected %q, got %q", "quick-fox", name)
		}
	}
}
