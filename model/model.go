package model

type Data struct {
	GithubHandle, RepoName, Token string
	DeploymentId                  []string
	DeleteAll                     bool
}

type Deployment struct {
	ID          int    `json:"id"`
	Environment string `json:"original_environment"`
	CreatedAt   string `json:"created_at"`
}
