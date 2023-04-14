package utility

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Abhishek-Dobliyal/deploy-drop/model"
)

func dropDeployment(client http.Client, authHeader, deploymentUrl string) error {
	req, err := http.NewRequest("DELETE", deploymentUrl, nil)
	if err != nil {
		return fmt.Errorf("error creating deletion request. error: %s", err.Error())
	}
	req.Header.Add("Authorization", authHeader)

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error deleting deployment. error: %s", err.Error())
	}
	defer resp.Body.Close()

	return nil
}

func contains(deployment []model.Deployment, target int) bool {
	for _, data := range deployment {
		if data.ID == target {
			return true
		}
	}
	return false
}

func SendRequest(data model.Data) ([]*model.Deployment, error) {
	var (
		deployments []model.Deployment
		client      http.Client
	)
	allDeploymentsUrl := fmt.Sprintf("https://api.github.com/repos/%s/%s/deployments", data.GithubHandle, data.RepoName)
	authHeader := fmt.Sprintf("bearer %s", data.Token)

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

	if len(deployments) == 0 {
		return nil, fmt.Errorf("error locating deployments. no deployments found")
	}

	if data.DeploymentId == nil {
		for _, deployment := range deployments {
			deploymentUrl := fmt.Sprintf("%s/%d", allDeploymentsUrl, deployment.ID)
			err := dropDeployment(client, authHeader, deploymentUrl)
			if err != nil {
				fmt.Println(err.Error())
			}

		}
	} else {
		for _, deploymentId := range data.DeploymentId {
			if contains(deployments, deploymentId) {
				deploymentUrl := fmt.Sprintf("%s/%d", allDeploymentsUrl, deploymentId)
				err := dropDeployment(client, authHeader, deploymentUrl)
				if err != nil {
					fmt.Println(err.Error())
				}
			}
		}
	}
	return nil, nil
}
