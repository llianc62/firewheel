package email_test

/*
import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"strings"
	"testing"
	"time"

	jemail "github.com/jordan-wright/email"
	"github.com/stretchr/testify/assert"

	fw "github.com/LiangXianSen/firewheel"
	"github.com/LiangXianSen/firewheel/sender/email"
)

func TestEmailSendText(t *testing.T) {
	var err error
	var sender fw.Sender

	must := assert.New(t)
	sender, err = email.NewSender(
		"smtp.gmail.com:587",
		email.LoginAuth("username", "password"),
	)
	must.Nil(err)

	err = sender.Send(
		strings.NewReader("test message"),
		email.TextContentType(),
		email.Subject("email test"),
		email.From("test@gmail.com"),
		email.To("test@gmail.com", "example@gmail.com"),
		email.Bcc("example@gmail.com"),
		email.Cc("example@gmail.com"),
	)
	must.Nil(err)
}

func TestEmailSendHTML(t *testing.T) {
	var err error
	var sender fw.Sender

	must := assert.New(t)
	sender, err = email.NewSender(
		"smtp.gmail.com:587",
		email.LoginAuth("username", "password"),
	)
	must.Nil(err)

	err = sender.Send(
		strings.NewReader("<h1>Fancy HTML is supported, too!</h1>"),
		email.HTMLContentType(),
		email.Subject("email test"),
		email.From("test@gmail.com"),
		email.To("test@gmail.com", "example@gmail.com"),
		email.Bcc("example@gmail.com"),
		email.Cc("example@gmail.com"),
	)
	must.Nil(err)
}

func TestBuiltInSMTP(t *testing.T) {
	must := assert.New(t)
	conn, err := smtp.Dial("smtp.gmail.com:587")
	must.Nil(err)
	defer conn.Close()

	// Check connection to server is okay or not.
	err = conn.Noop()
	must.Nil(err)

	err = conn.StartTLS(&tls.Config{ServerName: "smtp.gmail.com"})
	must.Nil(err)

	auth := email.LoginAuth("username", "password")
	ok, param := conn.Extension("Auth")
	must.Equal(true, ok)
	fmt.Printf("SMTP Extension param: %s\n", param)
	err = conn.Auth(auth)
	must.Nil(err)

	// err = conn.Verify("test@gmail.com")
	// must.Nil(err)

	err = conn.Mail("test@gmail.com")
	must.Nil(err)
	err = conn.Rcpt("test@gmail.com")
	must.Nil(err)

	buf, err := conn.Data()
	must.Nil(err)
	_, err = fmt.Fprint(buf, "test to my self")
	must.Nil(err)
	err = buf.Close()
	must.Nil(err)

	err = conn.Quit()
	must.Nil(err)
}

func TestJordanEmailWithText(t *testing.T) {
	auth := email.LoginAuth("username", "password")

	e := jemail.NewEmail()
	e.From = "test@gmail.com"
	e.To = []string{"example@gmail.com"}
	e.Bcc = []string{"example@gmail.com"}
	e.Cc = []string{"example@gmail.com"}
	e.Subject = "email test"
	e.Text = []byte("hello test text")
	e.Send("smtp.gmail.com:587", auth)
}

func TestJordanEmailPool(t *testing.T) {
	must := assert.New(t)
	auth := email.LoginAuth("username", "password")
	pool, err := jemail.NewPool(
		"smtp.gmail.com:587",
		4,
		auth,
	)
	must.Nil(err)

	for i := 0; i < 10; i++ {
		e := jemail.NewEmail()
		e.From = "test@gmail.com"
		e.To = []string{"example@gmail.com"}
		e.Bcc = []string{"example@gmail.com"}
		e.Cc = []string{"example@gmail.com"}
		e.Subject = "email test"
		e.Text = []byte(fmt.Sprintf("hello test %d text", i))
		pool.Send(e, time.Second*1)
	}
}
*/
