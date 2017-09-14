package main

import (
	"time"

	"github.com/nlopes/slack"
)

func timeManagement(ev *slack.MessageEvent, rtm *slack.RTM) {
	//出勤処理
	if ev.Text == "出勤" {
		startTime := time.Now()
		//TODO 出勤処理
		rtm.SendMessage(rtm.NewOutgoingMessage(startTime.Format("2006/01/02 Mon 15:04:05")+"に出勤を確認しました", ev.Channel))
	}

	//退勤処理
	if ev.Text == "退勤" {
		endTime := time.Now()
		//TODO 退勤処理
		rtm.SendMessage(rtm.NewOutgoingMessage(endTime.Format("2006/01/02 Mon 15:04:05")+"に退勤を確認しました", ev.Channel))
	}

}
