package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "deploy-drop [command] [...args]",
		Short: "A CLI for automatic removal of deployments on Github.",
		Long: `Remove deployments of unused projects on Github. Uses official Github API & it's token to accomplish the task. `,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
