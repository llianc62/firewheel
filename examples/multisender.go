package main

import (
	"io"
	"log"
	"strings"

	fw "github.com/LiangXianSen/firewheel"
	"github.com/LiangXianSen/firewheel/sender/dingding"
	"github.com/LiangXianSen/firewheel/sender/email"
)

// MSender implements Messager.
type MSender struct {
	senders []fw.Sender
}

func (ms *MSender) Setup(senders ...fw.Sender) {
	ms.senders = append(ms.senders, senders...)
}

func (ms *MSender) Send(message io.Reader, opt ...fw.Option) (err error) {
	for _, sender := range ms.senders {
		if err = sender.Send(message, opt...); err != nil {
			log.Printf("send message failed: %s", err)
		}
	}
	return
}

func main() {
	var holder fw.Messager
	holder = new(MSender)

	// Setup email sender
	var err error
	var sender fw.Sender
	if sender, err = email.NewSender(
		"smtp.gmail.com:587",
		email.LoginAuth("username", "password"),
		email.TextContentType(),
		email.Subject("email test"),
		email.From("test@gmail.com"),
		email.To("example@gmail.com"),
		email.Bcc("example@gmail.com"),
		email.Cc("example@gmail.com"),
	); err == nil {
		holder.Setup(sender)
	} else {
		log.Fatal(err)
	}

	// Setup DingTalk sender
	holder.Setup(dingding.NewSender(
		"https://oapi.dingtalk.com/robot/send?access_token=",
		dingding.MessageMode(dingding.TextMode),
		dingding.KeyWords("keywords"),
	))

	// Send message by Messager which setups two sender.
	//
	// As above we sets up all fixed value when news sender instance.
	// for sure, you can just news sender without any options, gives
	// options during sending message, both okay, depends real situation.
	//
	// In fixed messaging case, I do not have to used to change sending
	// configuration. I just want to simply use Send() to deliver message,
	// so I sets up all values when news sender at beginning. I sends message
	// by multiple senders from same data source, I implements a Messager which
	// setups two senders.
	if err = holder.Send(strings.NewReader("test message")); err != nil {
		log.Fatal(err)
	}

	// Also In some cases, we cannot gives a fixed value during news sender.
	// there are some functional Option, like email.SubjectFunc(f func() string)
	// which provides hook function f somewhat can solve your problem.
}
