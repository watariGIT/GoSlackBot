package main

import (
	"fmt"
	"github.com/nlopes/slack"
	"math/rand"
	"time"
)

var results = []string{
	"大吉",
	"末吉",
	"中吉",
	"吉",
	"小吉",
	"凶",
	"大凶",
	"うんこ",
	"ちんこ(ﾎﾟﾛﾝ",
}

//くじ引き関数
func lottery(ev *slack.MessageEvent, rtm *slack.RTM) {
	if ev.Text == "占い" {
		rand.Seed(time.Now().UnixNano())
		resultStr := results[rand.Intn(len(results))]
		rtm.SendMessage(rtm.NewOutgoingMessage(fmt.Sprint("占いの結果は... \n", resultStr), ev.Channel))
	}
}
