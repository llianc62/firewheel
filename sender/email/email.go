// Package email provides sending email ability.
package email

import (
	"io"
	"io/ioutil"
	"net/smtp"
	"time"

	jemail "github.com/jordan-wright/email"

	fw "github.com/LiangXianSen/firewheel"
)

const (
	defaultPoolSize = 4
	defaultTimeout  = time.Second * 1
)

// Email implements Sender which sends email.
type Email struct {
	opts   *options
	sender *jemail.Pool
}

// Send sends email which you have to provides io.Reader include message body,
// and few options which implements Option.
func (e *Email) Send(message io.Reader, opt ...fw.Option) (err error) {
	var content []byte
	if content, err = ioutil.ReadAll(message); err != nil {
		return
	}

	opts := new(options)
	*opts = *e.opts

	var roll fw.Option
	roll = opts
	for _, o := range opt {
		roll = o.Apply(roll)
	}

	mail := &jemail.Email{
		Subject: opts.subject,
		From:    opts.from,
		To:      opts.to,
		Bcc:     opts.bcc,
		Cc:      opts.cc,
	}
	if opts.attachment != nil {
		if _, err = mail.Attach(opts.attachment.fd, opts.attachment.filename, ""); err != nil {
			return
		}
	}
	if opts.timeout <= 0 {
		opts.timeout = defaultTimeout
	}
	switch opts.contentType {
	case "HTML":
		mail.HTML = content
	case "Text":
		mail.Text = content
	default:
		mail.Text = content
	}

	return e.sender.Send(mail, opts.timeout)
}

// NewSender returns Sender which sends email.
// requires SMTP server address, like "smtp.gmail.com:587",
// smtp.Auth which you have to implements and defines how get throug authorization.
func NewSender(smtpServer string, auth smtp.Auth, opt ...fw.Option) (sender fw.Sender, err error) {
	opts := new(options)
	var roll fw.Option
	roll = opts
	for _, o := range opt {
		roll = o.Apply(roll)
	}

	if opts.poolSize <= 0 {
		opts.poolSize = defaultPoolSize
	}

	var pool *jemail.Pool
	if pool, err = jemail.NewPool(
		smtpServer,
		opts.poolSize,
		auth,
		opts.tlsConf...,
	); err != nil {
		return
	}
	return &Email{
		opts:   opts,
		sender: pool,
	}, nil
}
