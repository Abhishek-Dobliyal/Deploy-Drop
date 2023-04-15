package cmd

import (
	"fmt"

	"github.com/Abhishek-Dobliyal/deploy-drop/model"
	"github.com/Abhishek-Dobliyal/deploy-drop/utility"

	"github.com/spf13/cobra"
)

var (
	githubHandleSearch string
	repoNameSearch     string

	search = &cobra.Command{
		Use:   "search",
		Short: "Searches for deployemnt(s) associated with the given repository.",
		Long:  "Searches for deployment(s) associated with the given repository.\nRequires Github handle & Repository name.",
		Run: func(cmd *cobra.Command, args []string) {
			data := model.Data{
				GithubHandle: &githubHandleSearch,
				RepoName:     &repoNameSearch,
			}

			deployments, err := utility.SearchDeployment(data)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			fmt.Println(deployments)
		},
	}
)

func init() {
	search.Flags().StringVarP(&githubHandleSearch, "handle", "u", "", "Github Repository Link (Required)")
	search.Flags().StringVarP(&repoNameDrop, "repo", "r", "", "Repository Name (Required)")

	search.MarkFlagRequired("handle")
	search.MarkFlagRequired("repo")

	rootCmd.AddCommand(search)
}
