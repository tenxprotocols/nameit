package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/tenxprotocols/nameit/lists"
)

// loadWordLists loads adjectives and nouns based on provided flags
func loadWordLists() ([]string, []string) {
	var adjectives, nouns []string
	var err error

	// Load adjectives
	if len(adjectivesList) > 0 {
		// Use provided adjectives list
		adjectives = adjectivesList
	} else if adjectivesFile != "" {
		// Load adjectives from file
		adjectives, err = loadWordsFromFile(adjectivesFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading adjectives file: %v\n", err)
			os.Exit(1)
		}
	} else {
		// Load default adjectives based on mode
		switch mode {
		case "modern":
			adjectives = lists.ModernAdjectivesList
		case "animal":
			adjectives = lists.AnimalAdjectivesList
		default:
			adjectives = lists.HerokuAdjectivesList
		}
	}

	// Load nouns
	if len(nounsList) > 0 {
		// Use provided nouns list
		nouns = nounsList
	} else if nounsFile != "" {
		// Load nouns from file
		nouns, err = loadWordsFromFile(nounsFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading nouns file: %v\n", err)
			os.Exit(1)
		}
	} else {
		// Load default nouns based on mode
		switch mode {
		case "modern":
			nouns = lists.ModernNounsList
		case "animal":
			nouns = lists.AnimalNounsList
		default:
			nouns = lists.HerokuNounsList
		}
	}

	return adjectives, nouns
}

// loadWordsFromFile loads words from a file, one per line
func loadWordsFromFile(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		if word != "" {
			words = append(words, word)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}
