package cmd

import (

	"github.com/Abhishek-Dobliyal/deploy-drop/utility"
	"github.com/Abhishek-Dobliyal/deploy-drop/model"

	"github.com/spf13/cobra"
)

const LONG_DESC = "Drops any deployment(s) associated with the given repository.\nPass Repository URL & Github Token as flags. Make sure the token is read_deployments authorized."

var (
	GithubHandle string
	RepoName     string
	Token        string
	DeploymentId []string

	drop = &cobra.Command{
		Use:   "drop",
		Short: "Drops any deployment(s) associated with the given repository.",
		Long:  LONG_DESC,
		Run: func(cmd *cobra.Command, args []string) {
			data := model.Data{
				GithubHandle: GithubHandle,
				RepoName:     RepoName,
				Token:        Token,
				DeploymentId: DeploymentId,
			}

			utility.SendRequest(data)
		},
	}
)

func init() {
	drop.Flags().StringVarP(&GithubHandle, "handle", "", "", "Github Repository Link (Required)")
	drop.Flags().StringVarP(&Token, "token", "t", "", "Github Token (Required, read_deployments authorized)")
	drop.Flags().StringVarP(&RepoName, "repo", "r", "", "Github Token (Required, read_deployments authorized)")
	drop.Flags().StringSliceVar(&DeploymentId, "ids", nil, "Github Token (Required, read_deployments authorized)")

	drop.MarkFlagRequired("url")
	drop.MarkFlagRequired("token")
	drop.MarkFlagRequired("repo")

	rootCmd.AddCommand(drop)
}
