package dingding_test

import (
	"log"
	"strings"

	fw "github.com/LiangXianSen/firewheel"
	"github.com/LiangXianSen/firewheel/sender/dingding"
)

func ExampleDingDing_Send_text() {
	var err error
	var sender fw.Sender

	webhook := "https://oapi.dingtalk.com/robot/send?access_token="
	sender = dingding.NewSender(webhook)
	if err = sender.Send(strings.NewReader("keywords: test message")); err != nil {
		log.Fatal(err)
	}
}

func ExampleDingDing_Send_markdown() {
	var err error
	var sender fw.Sender

	keywords := "keywords"
	webhook := "https://oapi.dingtalk.com/robot/send?access_token="
	sender = dingding.NewSender(webhook)
	pattern := `
	## topic: license expired
	something happened, pls checkout.
	## detail:
	- one
	- two
	- three
	`

	if err = sender.Send(
		strings.NewReader(pattern),
		dingding.MessageMode(dingding.MarkdownMode),
		dingding.KeyWords(keywords),
	); err != nil {
		log.Fatal(err)
	}
}

func ExampleDingDing_Send_link() {
	var err error
	var sender fw.Sender

	keywords := "keywords"
	webhook := "https://oapi.dingtalk.com/robot/send?access_token="
	sender = dingding.NewSender(webhook)
	if err = sender.Send(
		strings.NewReader("test message"),
		dingding.MessageMode(dingding.LinkMode),
		dingding.KeyWords(keywords),
		dingding.MessageURL("https://gist.github.com/asim/d3dea1832609b32538838529e693f9bf"),
		dingding.PicURL("https://ss0.bdstatic.com/70cFvHSh_Q1YnxGkpoWK1HF6hhy/it/u=110199744,2712833928&fm=26&gp=0.jpg"),
	); err != nil {
		log.Fatal(err)
	}
}

func ExampleDingDing_Send_actionCard() {
	var err error
	var sender fw.Sender

	keywords := "keywords"
	webhook := "https://oapi.dingtalk.com/robot/send?access_token="
	sender = dingding.NewSender(webhook)
	if err = sender.Send(
		strings.NewReader("test message"),
		dingding.MessageMode(dingding.ActionCardMode),
		dingding.KeyWords(keywords),
		dingding.SingleTitle("read more"),
		dingding.SingleURL("https://gist.github.com/asim/d3dea1832609b32538838529e693f9bf"),
	); err != nil {
		log.Fatal(err)
	}
}
