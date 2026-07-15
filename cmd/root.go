// Package cmd implements the command-line interface for the nameit tool
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Config file path
	cfgFile string

	// Name generation options
	mode           string   // Generation style: "modern", "heroku", or "animal"
	adjectivesList []string // Custom list of adjectives
	nounsList      []string // Custom list of nouns
	adjectivesFile string   // Path to file with custom adjectives
	nounsFile      string   // Path to file with custom nouns

	// Formatting options
	prefix       string // Prepend a prefix
	separator    string // Character(s) between words
	appendRandom bool   // Append a random suffix
	randomChars  string // Characters for random suffix
	randomLength int    // Length of the random suffix
	count        int    // Number of names to generate
	outputFormat string // Output format (text, json, yaml)
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "nameit",
	Short: "Generate memorable random names",
	Long: `A CLI application to generate Heroku-like memorable random names.
It allows you to customize the word list, separator, and append a number.
Output formats include plain text (one per line), JSON array, or YAML array.

Examples:
  # Generate a single Heroku-style name
  nameit --mode=heroku

  # Generate an animal-themed name
  nameit --mode=animal

  # Generate 5 modern-style names with custom separator
  nameit --mode=modern --count=5 --separator="_"

  # Generate names using custom word lists
  nameit --adjectives-list=red,blue,green --nouns-list=apple,banana,orange

  # Generate names using words from files
  nameit --adjectives-file=./my-adjectives.txt --nouns-file=./my-nouns.txt

  # Output as JSON
  nameit --count=10 --output=json`,
	// Run the main command logic
	Run: func(cmd *cobra.Command, args []string) {
		applyConfig()
		adjectives, nouns := loadWordLists() // Load adjectives and nouns based on provided flags
		names := generateNames(adjectives, nouns, prefix, separator, appendRandom, randomChars, randomLength, count)
		outputNames(names)
	},
}

// SetVersionInfo sets the version string shown by `nameit --version`.
// Called from main with values injected at build time.
func SetVersionInfo(version, commit, date string) {
	rootCmd.Version = fmt.Sprintf("%s (commit %s, built %s)", version, commit, date)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Global flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.nameit.yaml)")

	// Mode selection
	rootCmd.Flags().StringVar(&mode, "mode", "modern", "Generation mode (\"modern\", \"heroku\", or \"animal\")")

	// Generation options
	rootCmd.Flags().IntVar(&count, "count", 1, "Number of names to generate")
	rootCmd.Flags().StringVar(&prefix, "prefix", "", "Prepend prefix to name")
	rootCmd.Flags().StringVar(&separator, "separator", "-", "Separator between words")
	rootCmd.Flags().StringVar(&outputFormat, "output", "text", "Output format: text, json, or yaml")
	rootCmd.Flags().BoolVar(&appendRandom, "append-random", false, "Whether to append a random token to the end of the name")
	rootCmd.Flags().StringVar(&randomChars, "random-chars", "0123456789", "Characters to use when generating the random token")
	rootCmd.Flags().IntVar(&randomLength, "random-length", 3, "Length of the random token")

	// Word list options
	rootCmd.Flags().StringSliceVar(&adjectivesList, "adjectives-list", []string{}, "List of adjectives to use in name")
	rootCmd.Flags().StringSliceVar(&nounsList, "nouns-list", []string{}, "List of nouns to use in name")
	rootCmd.Flags().StringVar(&adjectivesFile, "adjectives-file", "", "Path to file containing adjectives, one per line")
	rootCmd.Flags().StringVar(&nounsFile, "nouns-file", "", "Path to file containing nouns, one per line")

	// Mark flags as mutually exclusive
	rootCmd.MarkFlagsMutuallyExclusive("mode", "adjectives-list")
	rootCmd.MarkFlagsMutuallyExclusive("mode", "adjectives-file")
	rootCmd.MarkFlagsMutuallyExclusive("mode", "nouns-list")
	rootCmd.MarkFlagsMutuallyExclusive("mode", "nouns-file")
	rootCmd.MarkFlagsMutuallyExclusive("adjectives-list", "adjectives-file")
	rootCmd.MarkFlagsMutuallyExclusive("nouns-list", "nouns-file")

	// Bind flags to viper for configuration file support
	viper.BindPFlag("mode", rootCmd.Flags().Lookup("mode"))

	viper.BindPFlag("count", rootCmd.Flags().Lookup("count"))
	viper.BindPFlag("output", rootCmd.Flags().Lookup("output"))
	viper.BindPFlag("prefix", rootCmd.Flags().Lookup("prefix"))
	viper.BindPFlag("separator", rootCmd.Flags().Lookup("separator"))
	viper.BindPFlag("append-random", rootCmd.Flags().Lookup("append-random"))
	viper.BindPFlag("random-chars", rootCmd.Flags().Lookup("random-chars"))
	viper.BindPFlag("random-length", rootCmd.Flags().Lookup("random-length"))

	viper.BindPFlag("adjectives-list", rootCmd.Flags().Lookup("adjectives-list"))
	viper.BindPFlag("nouns-list", rootCmd.Flags().Lookup("nouns-list"))
	viper.BindPFlag("adjectives-file", rootCmd.Flags().Lookup("adjectives-file"))
	viper.BindPFlag("nouns-file", rootCmd.Flags().Lookup("nouns-file"))
}

// applyConfig resolves each option through viper so values from the config
// file and environment take effect, while flags set on the command line keep
// precedence (viper returns the flag value when the flag was changed).
func applyConfig() {
	mode = viper.GetString("mode")
	count = viper.GetInt("count")
	outputFormat = viper.GetString("output")
	prefix = viper.GetString("prefix")
	separator = viper.GetString("separator")
	appendRandom = viper.GetBool("append-random")
	randomChars = viper.GetString("random-chars")
	randomLength = viper.GetInt("random-length")
	adjectivesList = viper.GetStringSlice("adjectives-list")
	nounsList = viper.GetStringSlice("nouns-list")
	adjectivesFile = viper.GetString("adjectives-file")
	nounsFile = viper.GetString("nouns-file")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".nameit" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".nameit")
	}

	viper.SetEnvPrefix("NAMEIT")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
