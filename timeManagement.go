package main

import (
	"encoding/csv"
	"fmt"
	"github.com/nlopes/slack"
	"log"
	"os"
	"time"
)

func timeManagement(ev *slack.MessageEvent, rtm *slack.RTM) {
	//出勤処理
	if ev.Text == "出勤" {
		inTime := time.Now()

		// 出勤書き込み
		file, err := os.OpenFile("files/timestamp_"+ev.User+".csv", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		//すでに出勤していないかの確認
		csvReader := csv.NewReader(file)
		records, _ := csvReader.ReadAll()

		for _, record := range records {
			if record[0] == inTime.Format("2006/01/02") {
				rtm.SendMessage(rtm.NewOutgoingMessage("すでに出勤しています", ev.Channel))
				return
			}
		}

		//出勤時間書き込み
		csvWriter := csv.NewWriter(file)
		newRecord := []string{inTime.Format("2006/01/02"), inTime.Format("15:04:05"), ""}
		csvWriter.Write(newRecord)
		csvWriter.Flush()
		rtm.SendMessage(rtm.NewOutgoingMessage(inTime.Format("2006/01/02 Mon 15:04:05")+"に出勤を確認しました", ev.Channel))
	}

	//退勤処理
	if ev.Text == "退勤" {
		outTime := time.Now()

		//退勤書き込み
		readfile, err := os.Open("files/timestamp_" + ev.User + ".csv")
		if err != nil {
			log.Fatal(err)
		}
		defer readfile.Close()

		//出勤したレコードの取得
		csvReader := csv.NewReader(readfile)
		records, _ := csvReader.ReadAll()
		for _, record := range records {
			if record[0] == outTime.Format("2006/01/02") {
				record[2] = outTime.Format("15:04:05")
				//勤務時間の計算+出力
				inTime, _ := time.Parse("2006/01/02 15:04:05 MST", record[0]+" "+record[1]+" JST")
				rtm.SendMessage(rtm.NewOutgoingMessage(outTime.Format("2006/01/02 15:04:05")+"に仕事がおわった!", ev.Channel))
				rtm.SendMessage(rtm.NewOutgoingMessage("今日は"+inTime.Format("15:04:05")+"に出勤だからぁ...", ev.Channel))
				subTime := outTime.Sub(inTime)
				rtm.SendMessage(rtm.NewOutgoingMessage(fmt.Sprintf("%d時間%d分 働きました!!", int(subTime.Hours())%24, int(subTime.Minutes())%60), ev.Channel))
			}
		}
		//TODO 出勤時間がない時の処理。
		//TODO すでに退勤時間ある時の変更確認処理

		//退勤時間書き込み
		writefile, err := os.Create("files/timestamp_" + ev.User + ".csv")
		if err != nil {
			log.Fatal(err)
		}
		defer writefile.Close()

		csvWriter := csv.NewWriter(writefile)
		csvWriter.WriteAll(records)
	}

	//勤怠管理機能のファイルアップロード
	if ev.Text == "タイムカード" {
		file,_:=os.Open("files/timestamp_" + ev.User + ".csv")
		defer file.Close()
				
		params := slack.FileUploadParameters{
			Title: "タイムカードです",
			File:  "files/timestamp_" + ev.User + ".csv",
			Channels : []string{ev.Channel},
		}
		
		_, err := rtm.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
	}
}
