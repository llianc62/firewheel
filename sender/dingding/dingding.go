// Package dingding provides sending message ability on DingTalk platform.
package dingding

import (
	"io"
	"io/ioutil"

	"github.com/royeo/dingrobot"

	fw "github.com/LiangXianSen/firewheel"
)

// DingDing implements Sender which binds with dingrobot.
type DingDing struct {
	opts   *options
	sender dingrobot.Roboter
}

// NewSender returns Sender which provides meesage delivery on DingTalk platform.
func NewSender(webhook string, opt ...fw.Option) fw.Sender {
	opts := new(options)
	var roll fw.Option
	roll = opts
	for _, o := range opt {
		roll = o.Apply(roll)
	}

	return &DingDing{
		opts:   opts,
		sender: dingrobot.NewRobot(webhook),
	}
}

// Send sends message which provides several format:
// Text, Markdown, Link, ActionCard, etc.
func (d *DingDing) Send(message io.Reader, opt ...fw.Option) (err error) {
	opts := new(options)
	*opts = *d.opts

	var roll fw.Option
	roll = opts
	for _, o := range opt {
		roll = o.Apply(roll)
	}

	switch opts.mode {
	case TextMode:
		return d.sendText(message, opts)
	case MarkdownMode:
		return d.sendMarkdown(message, opts)
	case LinkMode:
		return d.sendLink(message, opts)
	case ActionCardMode:
		return d.sendActionCard(message, opts)
	default:
		return d.sendText(message, opts)
	}
}

func (d *DingDing) sendText(message io.Reader, opts *options) (err error) {
	var content []byte
	if content, err = ioutil.ReadAll(message); err != nil {
		return
	}

	if err = d.sender.SendText(
		string(content),
		opts.atMobiles,
		opts.isAtAll,
	); err != nil {
		return
	}
	return
}

func (d *DingDing) sendMarkdown(message io.Reader, opts *options) (err error) {
	var content []byte
	if content, err = ioutil.ReadAll(message); err != nil {
		return
	}

	if err = d.sender.SendMarkdown(
		opts.keywords,
		string(content),
		opts.atMobiles,
		opts.isAtAll,
	); err != nil {
		return
	}
	return
}

func (d *DingDing) sendLink(message io.Reader, opts *options) (err error) {
	var content []byte
	if content, err = ioutil.ReadAll(message); err != nil {
		return
	}

	if err = d.sender.SendLink(
		opts.keywords,
		string(content),
		opts.messageURL,
		opts.picURL,
	); err != nil {
		return
	}
	return
}

func (d *DingDing) sendActionCard(message io.Reader, opts *options) (err error) {
	var content []byte
	if content, err = ioutil.ReadAll(message); err != nil {
		return
	}

	if err = d.sender.SendActionCard(
		opts.keywords,
		string(content),
		opts.singleTitle,
		opts.singleURL,
		opts.btnOrientation,
		opts.hideAvatar,
	); err != nil {
		return
	}
	return
}
