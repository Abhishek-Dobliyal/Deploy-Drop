package cmd

import (
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
			utility.DropDeployment(data)
		},
	}
)

func init() {
	rootCmd.AddCommand(search)
}
