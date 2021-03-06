package main

import (
	"fmt"
	"os"

	"github.com/nlopes/slack"
)

func main() {

	slackToken := os.Getenv("SLACK_TOKEN")

	api := slack.New(slackToken)
	// If you set debugging, it will log all requests to the console
	// Useful when encountering issues
	// api.SetDebug(true)
	groups, err := api.GetGroups(false)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	for _, group := range groups {
		fmt.Printf("ID: %s, Name: %s\n", group.ID, group.Name)
	}
}
