package dingding

import (
	fw "github.com/LiangXianSen/firewheel"
)

// DingTalk group robot api doc:
// https://developers.dingtalk.com/document/app/develop-enterprise-internal-robots

// MsgMode indicates message format.
type MsgMode int

const (
	// TextMode indicates text message you will send.
	TextMode MsgMode = iota + 1

	// MarkdownMode indicates markdown message you will send.
	MarkdownMode

	// LinkMode indicates link message you will send.
	LinkMode

	// ActionCardMode indicates acction card message you will send.
	ActionCardMode
)

// options implements Option.
type options struct {
	keywords       string
	atMobiles      []string
	isAtAll        bool
	messageURL     string
	picURL         string
	singleTitle    string
	singleURL      string
	btnOrientation string
	hideAvatar     string
	mode           MsgMode

	f func(*options)
}

func (opts *options) Apply(opt fw.Option) fw.Option {
	opts.f(opt.(*options))
	return opt
}

// KeyWords set keywords.
func KeyWords(keywords string) fw.Option {
	return &options{
		f: func(o *options) {
			o.keywords = keywords
		},
	}
}

// MessageMode set a MsgMode, like: TextMode, MarkdownMode, LinkMode and ActionCardMode.
func MessageMode(mode MsgMode) fw.Option {
	return &options{
		f: func(o *options) {
			o.mode = mode
		},
	}
}

// AtMobiles sets atMobiles.
func AtMobiles(atMobiles ...string) fw.Option {
	return &options{
		f: func(o *options) {
			o.atMobiles = append(o.atMobiles, atMobiles...)
		},
	}
}

// IsAtAll sets isAtAll.
func IsAtAll() fw.Option {
	return &options{
		f: func(o *options) {
			o.isAtAll = true
		},
	}
}

// MessageURL sets messageURL.
func MessageURL(messageURL string) fw.Option {
	return &options{
		f: func(o *options) {
			o.messageURL = messageURL
		},
	}
}

// PicURL sets picURL.
func PicURL(picURL string) fw.Option {
	return &options{
		f: func(o *options) {
			o.picURL = picURL
		},
	}
}

// SingleTitle sets singleTitle.
func SingleTitle(singleTitle string) fw.Option {
	return &options{
		f: func(o *options) {
			o.singleTitle = singleTitle
		},
	}
}

// SingleURL sets singleURL.
func SingleURL(singleURL string) fw.Option {
	return &options{
		f: func(o *options) {
			o.singleURL = singleURL
		},
	}
}

// BtnOrientation sets btnOrientation.
func BtnOrientation(btnOrientation string) fw.Option {
	return &options{
		f: func(o *options) {
			o.btnOrientation = btnOrientation
		},
	}
}

// HideAvatar sets hideAvatar.
func HideAvatar(hideAvatar string) fw.Option {
	return &options{
		f: func(o *options) {
			o.hideAvatar = hideAvatar
		},
	}
}
