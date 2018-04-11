package main

import (
	"os"
	"net/http"
	"fmt"
	"github.com/nlopes/slack"
	"github.com/JinOketani/slack_image_bot/cli"
)

func server(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "success")
	token := slack.New(os.Getenv("SLACK_API_TOKEN"))
	cli.Run(token)
}

func main() {
	http.HandleFunc("/", server)
	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}
