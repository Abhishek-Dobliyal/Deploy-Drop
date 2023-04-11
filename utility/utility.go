package utility

import (
	"net/http"
	"net/url"
)

func ParseURL(repoUrl string) (*string, error) {
	u, err := url.ParseRequestURI(repoUrl)
	if err != nil {
		return nil, err
	}

	uString := u.String()
	return &uString, nil
}

func SendRequest(parsedUrl, method string) (*http.Response, error) {
	resp, err := http.Post(parsedUrl, method, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
