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
			fmt.Printf("\n%s\n", Yellow("The following deployments were dropped:"))
			Printer.Print(dropped)
		},
	}
)

func init() {
	drop.Flags().StringVarP(&githubHandleDrop, "handle", "u", "", fmt.Sprintf("%s %s", Cyan("Github Handle"), Red("(Required)")))
	drop.Flags().StringVarP(&token, "token", "t", "", fmt.Sprintf("%s %s %s", Cyan("Github Token (Required"), Yellow("read_deployments"), Cyan("authorized)")))
	drop.Flags().StringVarP(&repoNameDrop, "repo", "r", "", fmt.Sprintf("%s %s", Cyan("Repository Name"), Red("(Required)")))
	drop.Flags().IntSliceVarP(&deploymendIdDrop, "ids", "i", nil, fmt.Sprintf("%s %s", Cyan("Deployment Id of the deployment to drop,"), Red("(Optional, If not specified all the associated deployments will be dropped)")))

	drop.MarkFlagRequired("url")
	drop.MarkFlagRequired("token")
	drop.MarkFlagRequired("repo")

	rootCmd.AddCommand(drop)
}
