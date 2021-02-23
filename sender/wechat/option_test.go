package wechat

import (
	"testing"

	fw "github.com/LiangXianSen/firewheel"
	"github.com/stretchr/testify/assert"
)

func TestMessageType(t *testing.T) {
	must := assert.New(t)

	opts := new(options)
	var roll fw.Option
	roll = opts
	MessageType(TextType).Apply(roll)
	must.Equal("text", opts.message.MessageType)
}

func TestCue(t *testing.T) {
	must := assert.New(t)

	opts := new(options)
	var roll fw.Option
	roll = opts
	Cue("@all").Apply(roll)
	must.Equal("@all", opts.message.Text.MentionList[0])
}

func TestCuePhone(t *testing.T) {
	must := assert.New(t)

	opts := new(options)
	var roll fw.Option
	roll = opts
	Cue("18512345678").Apply(roll)
	must.Equal("18512345678", opts.message.Text.MentionMobileList[0])
}

func TestArticle(t *testing.T) {
	must := assert.New(t)

	opts := new(options)
	var roll fw.Option
	roll = opts
	Article("test", "test message", "http://test.com", "http://test.com/1.png").Apply(roll)
	must.Equal("test", opts.message.News.Articles[0].Title)
	must.Equal("test message", opts.message.News.Articles[0].Description)
	must.Equal("http://test.com", opts.message.News.Articles[0].URL)
	must.Equal("http://test.com/1.png", opts.message.News.Articles[0].PicURL)
}

func TestFilename(t *testing.T) {
	must := assert.New(t)

	opts := new(options)
	var roll fw.Option
	roll = opts
	Filename("test.txt").Apply(roll)
	must.Equal("test.txt", opts.filename)
}
