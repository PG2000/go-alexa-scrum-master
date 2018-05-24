package jira

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
	"net/http"
)

type JiraClient struct {
	*jira.Client
}

func (client *JiraClient) IssueExists(issueId string) bool {
	fmt.Printf("Requesting issue with id: %s\n", issueId)
	_, response, e := client.Issue.Get(issueId, nil)

	if e != nil {
		fmt.Printf("An error occured while fetching the issue with id: %s, error: %+v\n", issueId, e)
	}

	return response.StatusCode == 200

}

func CreateClient(config *JiraConfig) *JiraClient {

	var httpClient *http.Client
	if (len(config.JiraUsername) > 0) {
		basicAuth := jira.BasicAuthTransport{
			Username: config.JiraUsername,
			Password: config.JiraPassword,
		}

		httpClient = basicAuth.Client()
	}
	cliente, e := jira.NewClient(httpClient, config.JiraEndpoint)

	if e != nil {
		fmt.Printf("Error occured while creating client %s\n", e)
	}

	return &JiraClient{
		cliente,
	}

}
