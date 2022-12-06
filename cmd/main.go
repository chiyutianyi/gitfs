package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/chiyutianyi/gitfs/pkg/version"
)

// Cmd represents the base command when called without any subcommands
var Cmd = &cobra.Command{
	Use:     "gitfs",
	Short:   "Git visual file system",
	Version: version.Version(),
}

func main() {
	if err := Cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "start git-worktree error: %v", err)
		os.Exit(1)
	}
}
