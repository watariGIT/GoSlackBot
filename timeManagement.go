package main

import (
	"github.com/nlopes/slack"
	"time"
	"log"
	"os"
	"fmt"
	"bufio"
	"strings"
)

func timeManagement(ev *slack.MessageEvent, rtm *slack.RTM) {
	//出勤処理
	if ev.Text == "出勤" {
		inTime := time.Now()

		// 出勤書き込み
		file, err := os.OpenFile("files/timestamp_"+ev.User+".txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		//すでに出勤していないかの確認
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			scanStr:=scanner.Text()
			if strings.HasPrefix(scanStr,inTime.Format("2006/01/02")) && strings.HasSuffix(scanStr,"出勤"){
				rtm.SendMessage(rtm.NewOutgoingMessage("すでに出勤しています", ev.Channel))		
				return
			}
		}
		
		fmt.Fprintln(file,inTime.Format("2006/01/02,15:04:05")+",出勤" )
		rtm.SendMessage(rtm.NewOutgoingMessage(inTime.Format("2006/01/02 Mon 15:04:05")+"に出勤を確認しました", ev.Channel))
	}

	//退勤処理
	if ev.Text == "退勤" {
		outTime := time.Now()

		//退勤書き込み
		file, err := os.OpenFile("files/timestamp_"+ev.User+".txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		fmt.Fprintln(file,outTime.Format("2006/01/02,15:04:05")+",退勤" )
		
		rtm.SendMessage(rtm.NewOutgoingMessage(outTime.Format("2006/01/02 Mon 15:04:05")+"に退勤を確認しました", ev.Channel))
	}

}
