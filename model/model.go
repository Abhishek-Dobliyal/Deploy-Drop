package model

type Data struct {
	GithubHandle, RepoName, Token *string
	DeploymentId                  []int
}

type Deployment struct {
	ID          int    `json:"id" header:"id"`
	Environment string `json:"original_environment" header:"environment"`
	CreatedAt   string `json:"created_at" header:"created_at"`
}
