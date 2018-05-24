package main

import (
	"fmt"
	client "go-scrum-master/jira"
	"flag"
)

func main() {
	config := client.GetConfig()
	jiraClient := client.CreateClient(&config)

	issue := flag.String("issue", "", "a string")

	flag.Parse()

	exists := jiraClient.IssueExists(*issue)
	fmt.Printf("Issue exists %t\n", exists)
}
