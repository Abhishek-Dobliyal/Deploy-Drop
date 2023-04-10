package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	URL  string
	drop = &cobra.Command{
		Use:   "drop <Repo URL>",
		Short: "Drops any deployment(s) associated with the given repository.",
		Long:  "Drops any deployment(s) associated with the given repository.",
		Run: func(cmd *cobra.Command, args []string) {
			val := cmd.Flag("url").Value
			fmt.Println(val)
		},
	}
)

func init() {
	drop.Flags().StringVarP(&URL, "url", "u", "", "Github Repository Link (Required)")
	drop.MarkFlagRequired("url")
	rootCmd.AddCommand(drop)
}
