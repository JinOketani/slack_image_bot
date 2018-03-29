package cli

import (
	"log"
	"regexp"
	"github.com/nlopes/slack"
	"strings"
)

func Run(api *slack.Client) int {
	var image string

	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.HelloEvent:
				log.Println("====Start====")
			case *slack.MessageEvent:
				if strings.Contains(ev.Text, "your user id") { // ex: <@U9U7Z06CV>
					rep := regexp.MustCompile(`your user id`) // ex: <@U9U7Z06CV>
					image = rep.ReplaceAllString(ev.Text, "")
					image = SearchImage(image)
					rtm.SendMessage(rtm.NewOutgoingMessage(image, ev.Channel))
				}
			case *slack.InvalidAuthEvent:
				log.Println("Invalid credentials")
				return 1
			}
		}
	}
}
