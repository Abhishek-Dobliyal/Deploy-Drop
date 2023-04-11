package cmd

import (
	"fmt"

	"github.com/Abhishek-Dobliyal/deploy-drop/utility"
	"github.com/spf13/cobra"
)

const LONG_DESC = "Drops any deployment(s) associated with the given repository.\nPass Repository URL & Github Token as flags. Make sure the token is read_deployments authorized."

var (
	URL   string
	Token string
	drop  = &cobra.Command{
		Use:   "drop",
		Short: "Drops any deployment(s) associated with the given repository.",
		Long:  LONG_DESC,
		Run: func(cmd *cobra.Command, args []string) {
			repoUrl, token := cmd.Flag("url").Value.String(), cmd.Flag("token").Value.String()
			parsedUrl, err := utility.ParseURL(repoUrl)
			if err != nil {
				fmt.Println("Error:", err.Error())
				fmt.Println(cmd.Usage())
				return
			}

			resp, err := utility.SendRequest(*parsedUrl, "POST")
			if err != nil {
				fmt.Println("Error:", err.Error())
				fmt.Println(cmd.Usage())
				return
			}

			fmt.Print(resp)
		},
	}
)

func init() {
	drop.Flags().StringVarP(&URL, "url", "u", "", "Github Repository Link (Required)")
	drop.Flags().StringVarP(&Token, "token", "t", "", "Github Token (Required, read_deployments authorized)")
	drop.MarkFlagRequired("url")
	drop.MarkFlagRequired("token")
	rootCmd.AddCommand(drop)
}
