package cmd

import (
	"fmt"

	"github.com/Abhishek-Dobliyal/deploy-drop/model"
	"github.com/Abhishek-Dobliyal/deploy-drop/utility"

	"github.com/spf13/cobra"
)

var (
	githubHandleDrop string
	repoNameDrop     string
	token            string
	deploymendIdDrop []int

	drop = &cobra.Command{
		Use:   "drop",
		Short: "Drops any deployment(s) associated with the given repository.",
		Long:  "Drops any deployment(s) associated with the given repository.\nRequires Github handle, Repository name & Github token as flags. Make sure the token is read_deployments authorized.",
		Run: func(cmd *cobra.Command, args []string) {
			data := model.Data{
				GithubHandle: &githubHandleDrop,
				RepoName:     &repoNameDrop,
				Token:        &token,
				DeploymentId: deploymendIdDrop,
			}
			dropped, err := utility.DropDeployment(data)
			if err != nil {
				fmt.Println(err.Error())
			}

			fmt.Println(dropped)
		},
	}
)

func init() {
	drop.Flags().StringVarP(&githubHandleDrop, "handle", "u", "", "Github Repository Link (Required)")
	drop.Flags().StringVarP(&token, "token", "t", "", "Github Token (Required, read_deployments authorized)")
	drop.Flags().StringVarP(&repoNameDrop, "repo", "r", "", "Repository Name (Required)")
	drop.Flags().IntSliceVarP(&deploymendIdDrop, "ids", "i", nil, "Deployment Id of the deployment to drop, (Optional, If not specified all the associated deployments will be dropped)")

	drop.MarkFlagRequired("url")
	drop.MarkFlagRequired("token")
	drop.MarkFlagRequired("repo")

	rootCmd.AddCommand(drop)
}
