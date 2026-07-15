package lists

import "testing"

func TestParseWordFile(t *testing.T) {
	words := parseWordFile("alpha\nbeta\n\ngamma\n")
	want := []string{"alpha", "beta", "gamma"}
	if len(words) != len(want) {
		t.Fatalf("expected %d words, got %d: %v", len(want), len(words), words)
	}
	for i, w := range want {
		if words[i] != w {
			t.Errorf("words[%d] = %q, want %q", i, words[i], w)
		}
	}
}

func TestEmbeddedListsNotEmpty(t *testing.T) {
	tests := []struct {
		name string
		list []string
	}{
		{"heroku adjectives", HerokuAdjectivesList},
		{"heroku nouns", HerokuNounsList},
		{"modern adjectives", ModernAdjectivesList},
		{"modern nouns", ModernNounsList},
		{"animal adjectives", AnimalAdjectivesList},
		{"animal nouns", AnimalNounsList},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if len(tt.list) == 0 {
				t.Errorf("%s list is empty", tt.name)
			}
			for _, w := range tt.list {
				if w == "" {
					t.Errorf("%s list contains an empty word", tt.name)
				}
			}
		})
	}
}
