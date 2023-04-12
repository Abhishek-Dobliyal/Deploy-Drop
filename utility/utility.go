package utility

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Abhishek-Dobliyal/deploy-drop/model"
)

func SendRequest(data model.Data) (*string, error) {
	var (
		deployments []model.Deployment
		client      http.Client
	)
	allDeploymentsUrl := fmt.Sprintf("https://api.github.com/repos/%s/%s/deployments", data.GithubHandle, data.RepoName)
	// authHeader := fmt.Sprintf("bearer %s", data.Token)

	req, err := http.NewRequest("GET", allDeploymentsUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request. Error: %s", err.Error())
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request. Error: %s", err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error response from server. Status code: %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&deployments)
	if err != nil {
		return nil, fmt.Errorf("error decoding response. Error: %s", err.Error())
	}
	defer resp.Body.Close()

	if data.DeploymentId == nil {
		for _, deployment := range deployments {
			deleteDeploymentUrl := fmt.Sprintf("%s/%d", allDeploymentsUrl, deployment.ID)
			fmt.Println(deleteDeploymentUrl)
		}
	}
	return nil, nil
}
