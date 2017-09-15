package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/nlopes/slack"
)

func todoList(ev *slack.MessageEvent, rtm *slack.RTM) {

	//メッセージが"add todo"始まりならTODOリストに追加
	if strings.HasPrefix(ev.Text, "add todo ") {
		todoStr := strings.Replace(ev.Text, "add todo ", "", 1)

		//useridのファイル名にTODOを追加
		file, err := os.OpenFile("files/todo_"+ev.User+".txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			log.Fatal(err)
			return
		}
		defer file.Close()
		fmt.Fprintln(file, todoStr)

		rtm.SendMessage(rtm.NewOutgoingMessage("「"+todoStr+"」をTODOに追加しました。", ev.Channel))
	}

	//メッセージが"ls todo"ならTODO一覧を表示
	if ev.Text == "ls todo" {
		//ファイルからTODOを読み取り
		file, err := os.Open("files/todo_" + ev.User + ".txt")

		if err != nil {
			log.Fatal(err)
			rtm.SendMessage(rtm.NewOutgoingMessage("ファイルの読み込みに失敗しました。", ev.Channel))
			return
		}
		defer file.Close()

		rtm.SendMessage(rtm.NewOutgoingMessage("やることの一覧です", ev.Channel))
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			rtm.SendMessage(rtm.NewOutgoingMessage("* "+scanner.Text(), ev.Channel))
		}
		rtm.SendMessage(rtm.NewOutgoingMessage("以上です。頑張ってください", ev.Channel))
	}
}
