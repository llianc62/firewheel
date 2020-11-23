package email_test

import (
	"log"
	"strings"

	fw "github.com/LiangXianSen/firewheel"
	"github.com/LiangXianSen/firewheel/sender/email"
)

func ExampleEmail_Send_text() {
	var err error
	var sender fw.Sender

	if sender, err = email.NewSender(
		"smtp.gmail.com:587",
		email.LoginAuth("username", "password"),
	); err != nil {
		log.Fatal(err)
	}

	if err = sender.Send(
		strings.NewReader("test message"),
		email.TextContentType(),
		email.Subject("email test"),
		email.From("test@gmail.com"),
		email.To("test@gmail.com", "example@gmail.com"),
		email.Bcc("example@gmail.com"),
		email.Cc("example@gmail.com"),
	); err != nil {
		log.Fatal(err)
	}
}

func ExampleEmail_Send_html() {
	var err error
	var sender fw.Sender

	if sender, err = email.NewSender(
		"smtp.gmail.com:587",
		email.LoginAuth("username", "password"),
	); err != nil {
		log.Fatal(err)
	}

	if err = sender.Send(
		strings.NewReader("<h1>Fancy HTML is supported, too!</h1>"),
		email.HTMLContentType(),
		email.Subject("email test"),
		email.From("test@gmail.com"),
		email.To("test@gmail.com", "example@gmail.com"),
		email.Bcc("example@gmail.com"),
		email.Cc("example@gmail.com"),
	); err != nil {
		log.Fatal(err)
	}
}
