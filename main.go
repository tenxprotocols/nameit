// Package main is the entry point for the nameit CLI tool,
// a name generator that creates memorable names with different styles.
package main

import "github.com/tenxprotocols/nameit/cmd"

// Version information injected at build time via ldflags.
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	cmd.SetVersionInfo(version, commit, date)
	cmd.Execute()
}
