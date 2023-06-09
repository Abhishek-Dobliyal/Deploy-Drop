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
			fmt.Printf("\n%s\n", Yellow("The following deployments were dropped:"))
			Printer.Print(deployments)
		},
	}
)

func init() {
	search.Flags().StringVarP(&githubHandleSearch, "handle", "u", "", fmt.Sprintf("%s %s", Cyan("Github Handle"), Red("(Required)")))
	search.Flags().StringVarP(&repoNameDrop, "repo", "r", "", fmt.Sprintf("%s %s", Cyan("Repository Name"), Red("(Required)")))

	search.MarkFlagRequired("handle")
	search.MarkFlagRequired("repo")

	rootCmd.AddCommand(search)
}
