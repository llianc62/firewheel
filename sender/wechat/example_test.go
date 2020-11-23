package wechat_test

import (
	"log"
	"os"
	"strings"

	fw "github.com/LiangXianSen/firewheel"
	"github.com/LiangXianSen/firewheel/sender/wechat"
)

func ExampleWeChat_Send_text() {
	var sender fw.Sender

	key := ""
	webhook := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send"
	sender = wechat.NewSender(webhook, key)
	if err := sender.Send(
		strings.NewReader("test message"),
		wechat.MessageType(wechat.TextType),
		wechat.Cue("@all"),
		wechat.Cue("18512345678"),
	); err != nil {
		log.Fatal(err)
	}
}

func ExampleWeChat_Send_markdown() {
	var sender fw.Sender

	key := ""
	webhook := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send"
	sender = wechat.NewSender(webhook, key)
	if err := sender.Send(
		strings.NewReader("# test message"),
		wechat.MessageType(wechat.MarkdownType),
	); err != nil {
		log.Fatal(err)
	}
}

func ExampleWeChat_Send_image() {
	fd, err := os.Open("../../testdata/timg.jpeg")
	if err != nil {
		log.Fatal(err)
	}

	var sender fw.Sender

	key := ""
	webhook := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send"
	sender = wechat.NewSender(webhook, key)
	if err = sender.Send(
		fd,
		wechat.MessageType(wechat.ImageType),
	); err != nil {
		log.Fatal(err)
	}
}

func ExampleWeChat_Send_news() {
	var sender fw.Sender

	key := ""
	webhook := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send"
	sender = wechat.NewSender(webhook, key)
	if err := sender.Send(
		strings.NewReader(""),
		wechat.MessageType(wechat.NewsType),
		wechat.Article("No.1", "Coding commnunity", "github.com", "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1606718584317&di=2546710d44d6582d9fe58bf18d2711a7&imgtype=0&src=http%3A%2F%2Foss.cyzone.cn%2F2016%2F0304%2F20160304022832427.png"),
		wechat.Article("No.2", "something else", "github.com", ""),
		wechat.Article("No.3", "last one", "github.com", ""),
	); err != nil {
		log.Fatal(err)
	}
}

func ExampleWeChat_Send_file() {
	fd, err := os.Open("../../testdata/report.csv")
	if err != nil {
		log.Fatal(err)
	}

	var sender fw.Sender

	key := ""
	webhook := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send"
	sender = wechat.NewSender(webhook, key)
	if err = sender.Send(
		fd,
		wechat.MessageType(wechat.FileType),
		wechat.Filename("report.csv"),
	); err != nil {
		log.Fatal(err)
	}
}
