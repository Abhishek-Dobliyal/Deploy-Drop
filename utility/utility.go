package utility

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Abhishek-Dobliyal/deploy-drop/model"
)

func SendRequest(data model.Data) (*string, error) {
	var deployments []model.Deployment

	allDeploymentsUrl := fmt.Sprintf("https://api.github.com/repos/%s/%s/deployments", data.GithubHandle, data.RepoName)
	// authHeader := fmt.Sprintf("bearer %s", data.Token)

	resp, err := http.Get(allDeploymentsUrl)
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
	return nil, nil
}
