package jira

import (
	"github.com/spf13/viper"
	"fmt"
)

type JiraConfig struct {
	JiraEndpoint        string;
	JiraUsername        string;
	JiraPassword        string;
	IssueExistsTemplate string;
}

var (
	config JiraConfig
)

func GetConfig() JiraConfig {

	viper.SetConfigFile("jira.yml")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println(err)
	}

	config := JiraConfig{
		JiraEndpoint:        viper.GetString("jira_endpoint"),
		JiraUsername:        viper.GetString("jira_username"),
		JiraPassword:        viper.GetString("jira_password"),
		IssueExistsTemplate: viper.GetString("issue_exists_template"),
	}
	return config
}
