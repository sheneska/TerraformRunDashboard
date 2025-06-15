package utils

import (
	"fmt"
	"io"
	"net/http"
	//"os"
	"strings"
)

var (
	apiBase = "https://app.terraform.io/api/v2"
	token   = "your-fake-token-or-leave-empty"
	org     = "your-org-name-or-leave-empty"
)

func buildRequest(method, url string, body *strings.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/vnd.api+json")
	return req, nil
}

func FetchWorkspaces() ([]byte, error) {
	url := fmt.Sprintf("%s/organizations/%s/workspaces", apiBase, org)
	req, err := buildRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return io.ReadAll(res.Body)
}

func FetchWorkspaceRuns(workspaceID string) ([]byte, error) {
	url := fmt.Sprintf("%s/workspaces/%s/runs", apiBase, workspaceID)
	req, err := buildRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return io.ReadAll(res.Body)
}

func TriggerRun(workspaceID string) ([]byte, error) {
	body := `{
		"data": {
			"type": "runs",
			"attributes": {
				"message": "Triggered from custom dashboard",
				"is-destroy": false
			},
			"relationships": {
				"workspace": {
					"data": {
						"type": "workspaces",
						"id": "` + workspaceID + `"
					}
				}
			}
		}
	}`
	url := fmt.Sprintf("%s/runs", apiBase)
	req, err := buildRequest("POST", url, strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return io.ReadAll(res.Body)
}

