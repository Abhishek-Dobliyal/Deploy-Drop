package cmd

import (
	"github.com/Abhishek-Dobliyal/deploy-drop/model"
	"github.com/Abhishek-Dobliyal/deploy-drop/utility"

	"github.com/spf13/cobra"
)

const LONG_DESC = "Drops any deployment(s) associated with the given repository.\nRequires Github handle, Repository name & Github token as flags. Make sure the token is read_deployments authorized."

var (
	GithubHandle string
	RepoName     string
	Token        string
	DeploymentId []int

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
	drop.Flags().StringVarP(&GithubHandle, "handle", "u", "", "Github Repository Link (Required)")
	drop.Flags().StringVarP(&Token, "token", "t", "", "Github Token (Required, read_deployments authorized)")
	drop.Flags().StringVarP(&RepoName, "repo", "r", "", "Repository Name (Required)")
	drop.Flags().IntSliceVarP(&DeploymentId, "ids", "i", nil, "Deployment Id of the deployment to drop, (Optional, If not specified all the associated deployments will be dropped)")

	drop.MarkFlagRequired("url")
	drop.MarkFlagRequired("token")
	drop.MarkFlagRequired("repo")

	rootCmd.AddCommand(drop)
}
