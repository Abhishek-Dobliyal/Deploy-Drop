package utility

import (
	"fmt"

	"github.com/Abhishek-Dobliyal/deploy-drop/model"
)

func SendRequest(data model.Data) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/deployments", data.GithubHandle, data.RepoName)
	authHeader := fmt.Sprintf("bearer %s", data.Token)

	fmt.Println(url, authHeader)
}
