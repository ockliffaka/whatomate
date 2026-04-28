package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/whatomate/whatomate/cmd"
)

var (
	// Version is set at build time via ldflags
	Version = "dev"
	// Commit is the git commit hash set at build time
	Commit = "none"
	// Date is the build date set at build time
	Date = "unknown"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "whatomate",
		Short: "WhatsApp automation tool",
		Long: `Whatomate is a CLI tool for automating WhatsApp messaging.
It allows you to send messages, manage contacts, and automate
WhatsApp workflows programmatically.`,
		Version: fmt.Sprintf("%s (commit: %s, built: %s)", Version, Commit, Date),
	}

	// Register subcommands
	rootCmd.AddCommand(cmd.NewSendCmd())
	rootCmd.AddCommand(cmd.NewServeCmd())
	rootCmd.AddCommand(cmd.NewVersionCmd(Version, Commit, Date))

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
