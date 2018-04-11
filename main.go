package main

import (
	"os"
	"net/http"
	"fmt"
	"github.com/nlopes/slack"
	"github.com/JinOketani/slack_image_bot/cli"
	"log"
)

func server(res http.ResponseWriter, req *http.Request) {
	token := slack.New(os.Getenv("SLACK_API_TOKEN"))
	os.Exit(cli.Run(token))
	fmt.Fprintln(res, "success")
}

func main() {
	http.HandleFunc("/", server)
	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Fatalf("error")
		panic(err)
	}
}