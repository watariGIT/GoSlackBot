package main

import (
	"log"
	"os"
	"strings"

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
				
				//メッセージが"add todo"始まりならTODOリストに追加
				if strings.HasPrefix(ev.Text, "add todo ") {
					todoStr := strings.Replace(ev.Text, "add todo ", "", 1)
					//TODO useridのファイル名にTODOを追加
					rtm.SendMessage(rtm.NewOutgoingMessage("「"+todoStr+"」をTODOに追加しました。", ev.Channel))
				}

				//メッセージが"ls todo"ならTODO一覧を表示
				if ev.Text == "ls todo" {
					//TODO useridのファイル名からTODOを取得
					//TODO 一覧をメッセージで表示
				}
				log.Printf("Message: %+v\n", ev)
				
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
