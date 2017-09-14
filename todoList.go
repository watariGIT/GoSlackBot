package main

import (
	"strings"

	"github.com/nlopes/slack"
)

func todoList(ev *slack.MessageEvent, rtm *slack.RTM) {
	//メッセージが"add todo"始まりならTODOリストに追加
	if strings.HasPrefix(ev.Text, "add todo ") {
		todoStr := strings.Replace(ev.Text, "add todo ", "", 1)
		//TODO useridのファイル名にTODOを追加
		rtm.SendMessage(rtm.NewOutgoingMessage("「"+todoStr+"」をTODOに追加しました。", ev.Channel))
	}

	//メッセージが"ls todo"ならTODO一覧を表示
	if ev.Text == "ls todo" {
		//TODO useridのファイル名からTODO
		//TODO 一覧をメッセージで表示
	}

}
