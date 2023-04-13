package utility

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Abhishek-Dobliyal/deploy-drop/model"
)

func dropDeployment(client http.Client, deploymentUrl string) {

}

func SendRequest(data model.Data) ([]*model.Deployment, error) {
	var (
		deployments []model.Deployment
		client      http.Client
	)
	allDeploymentsUrl := fmt.Sprintf("https://api.github.com/repos/%s/%s/deployments", data.GithubHandle, data.RepoName)
	// authHeader := fmt.Sprintf("bearer %s", data.Token)

	req, err := http.NewRequest("GET", allDeploymentsUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request. error: %s", err.Error())
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request. error: %s", err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error response from server. status code: %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&deployments)
	if err != nil {
		return nil, fmt.Errorf("error decoding response. error: %s", err.Error())
	}
	defer resp.Body.Close()

	if data.DeploymentId == nil {
		if len(deployments) == 0 {
			return nil, fmt.Errorf("error locating deployments. no deployments found")
		}
		for _, deployment := range deployments {
			deploymentUrl := fmt.Sprintf("%s/%d", allDeploymentsUrl, deployment.ID)
			dropDeployment(client, deploymentUrl)
		}
	}
	return nil, nil
}
