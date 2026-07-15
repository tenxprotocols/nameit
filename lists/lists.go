package lists

import (
	"slices"
	"strings"

	_ "embed"
)

func parseWordFile(file string) []string {
	parts := strings.Split(file, "\n") // Split the file on newline chars
	words := slices.DeleteFunc(parts, func(s string) bool { return s == "" })
	return words
}

//go:embed heroku/adjectives.txt
var herokuAdjectivesFile string
var HerokuAdjectivesList []string = parseWordFile(herokuAdjectivesFile)

//go:embed heroku/nouns.txt
var herokuNounsFile string
var HerokuNounsList []string = parseWordFile(herokuNounsFile)

//go:embed modern/adjectives.txt
var modernAdjectivesFile string
var ModernAdjectivesList []string = parseWordFile(modernAdjectivesFile)

//go:embed modern/nouns.txt
var modernNounsFile string
var ModernNounsList []string = parseWordFile(modernNounsFile)

//go:embed animal/adjectives.txt
var animalAdjectivesFile string
var AnimalAdjectivesList []string = parseWordFile(animalAdjectivesFile)

//go:embed animal/nouns.txt
var animalNounsFile string
var AnimalNounsList []string = parseWordFile(animalNounsFile)
