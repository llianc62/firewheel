package wechat

import (
	"strings"

	fw "github.com/LiangXianSen/firewheel"
)

// MsgType indicates message format.
type MsgType int

const (
	// TextType indicates text message you will send.
	TextType MsgType = iota + 1

	// MarkdownType indicates markdown message you will send.
	MarkdownType

	// ImageType indicates Image message you will send.
	ImageType

	// NewsType indicates news message you will send.
	NewsType

	// FileType indicates file message you will send.
	FileType
)

type options struct {
	key      string
	webhook  string
	filename string
	message  Message

	f func(*options)
}

func (opts *options) Apply(opt fw.Option) fw.Option {
	opts.f(opt.(*options))
	return opt
}

// MessageType sets MsgType, like: TextType, MarkdownType, ImageType, NewsType and FileType.
func MessageType(msgType MsgType) fw.Option {
	return &options{
		f: func(o *options) {
			switch msgType {
			case TextType:
				o.message.MessageType = "text"
			case MarkdownType:
				o.message.MessageType = "markdown"
			case ImageType:
				o.message.MessageType = "image"
			case NewsType:
				o.message.MessageType = "news"
			case FileType:
				o.message.MessageType = "file"
			}
		},
	}
}

// Cue sets mentions which could be phone number or user name.
func Cue(mention ...string) fw.Option {
	return &options{
		f: func(o *options) {
			if o.message.Text == nil {
				o.message.Text = new(textMessage)
			}
			for _, m := range mention {
				if strings.HasPrefix(m, "1") {
					o.message.Text.MentionMobileList = append(o.message.Text.MentionMobileList, m)
				} else {
					o.message.Text.MentionList = append(o.message.Text.MentionList, m)
				}
			}
		},
	}
}

// Article sets an article.
func Article(title, description, url, picurl string) fw.Option {
	return &options{
		f: func(o *options) {
			if o.message.News == nil {
				o.message.News = new(newsMessage)
			}
			o.message.News.Articles = append(o.message.News.Articles, article{
				Title:       title,
				Description: description,
				URL:         url,
				PicURL:      picurl,
			})
		},
	}
}

// Filename sets upload file name.
func Filename(filename string) fw.Option {
	return &options{
		f: func(o *options) {
			o.filename = filename
		},
	}
}
