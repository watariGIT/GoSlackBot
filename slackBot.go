package main

import (
	"log"
	"os"

	"github.com/nlopes/slack"
)

func run(api *slack.Client) int {
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.HelloEvent:
				log.Print("Hello Event")

			case *slack.MessageEvent:

				//TODOリストの処理
				todoList(ev, rtm)

				//勤怠管理の処理
				timeManagement(ev, rtm)

				//くじ引き
				lottery(ev, rtm)

			case *slack.InvalidAuthEvent:
				log.Print("Invalid credentials")
				return 1

			}
		}
	}
}

func main() {
	api := slack.New(getApiToken())
	os.Exit(run(api))
}
