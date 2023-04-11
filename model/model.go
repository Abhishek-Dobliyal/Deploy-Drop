package model

type Data struct {
	GithubHandle, RepoName, Token string
	DeploymentId                  []string
	DeleteAll                     bool
}
