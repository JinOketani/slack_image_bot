package main

import (
	"github.com/nlopes/slack"
	"os"
	"github.com/JinOketani/slack_image_bot/cli"
)

func main() {
	token := slack.New(os.Getenv("SLACK_API_TOKEN"))
	os.Exit(cli.Run(token))
}
