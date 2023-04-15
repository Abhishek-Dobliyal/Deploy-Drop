package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	Red     = color.New(color.FgRed, color.Bold).SprintfFunc()
	Cyan    = color.New(color.FgCyan, color.Bold).SprintfFunc()
	Yellow   = color.New(color.FgYellow, color.Italic).SprintfFunc()
	rootCmd = &cobra.Command{
		Use:   "deploy-drop",
		Short: "A CLI for automatic removal of deployments on Github.",
		Long:  `Remove deployments of unused projects on Github. Uses official Github API & it's token to accomplish the task. `,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
