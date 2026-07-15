# nameit

A command-line tool for generating memorable names with different styles.

## Description

`nameit` is a name generator CLI tool that creates memorable names by combining adjectives and nouns. It supports different naming styles and allows customization through various options.

## Features

- Generate names in different styles/modes:
  - Modern style
  - Heroku-like style
  - Animal style
- Customize with your own word lists
- Load custom adjectives and nouns from files
- Output as plain text, JSON, or YAML
- Configuration via config file or environment variables

## Installation

### Homebrew

```bash
brew install tenxprotocols/tap/nameit
```

### Go

```bash
go install github.com/tenxprotocols/nameit@latest
```

### Binaries

Pre-built binaries for Linux, macOS, and Windows are available on the
[releases page](https://github.com/tenxprotocols/nameit/releases).

## Usage

Basic usage:

```bash
nameit
```

With options:

```bash
# Generate name
nameit

# Pick a style
nameit --mode heroku
nameit --mode animal

# Specify a count
nameit --count 5

# Specify a prefix
nameit --prefix my-corp

# Specify a random suffix of 5 hex characters
nameit --append-random --random-length 5 --random-chars 0123456789abcdef

# Use custom word lists
nameit --adjectives-file /path/to/adjectives.txt --nouns-file /path/to/nouns.txt

# Output as JSON
nameit --count 10 --output json
```

## Configuration

`nameit` supports configuration via:

1. Command-line flags
2. Environment variables (prefixed with `NAMEIT_`, e.g. `NAMEIT_COUNT=5`)
3. Config file: `.nameit.yaml` in your home directory

## Releases

Releases are automated with [release-please](https://github.com/googleapis/release-please):
merges to `main` with [Conventional Commit](https://www.conventionalcommits.org/) messages
accumulate into a release PR; merging that PR tags a release, and
[GoReleaser](https://goreleaser.com) builds the binaries and updates the
[Homebrew tap](https://github.com/tenxprotocols/homebrew-tap).
