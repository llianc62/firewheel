package dingding

import (
	"testing"

	fw "github.com/LiangXianSen/firewheel"
	"github.com/stretchr/testify/assert"
)

func TestKeyWords(t *testing.T) {
	must := assert.New(t)

	opts := new(options)
	var roll fw.Option
	roll = opts
	KeyWords("test").Apply(roll)
	must.Equal("test", opts.keywords)
}

func TestMessageMode(t *testing.T) {
	must := assert.New(t)

	opts := new(options)
	var roll fw.Option
	roll = opts
	MessageMode(TextMode).Apply(roll)
	must.Equal(TextMode, opts.mode)
}

func TestAtMobiles(t *testing.T) {
	must := assert.New(t)

	opts := new(options)
	var roll fw.Option
	roll = opts
	AtMobiles("test").Apply(roll)
	must.Equal("test", opts.atMobiles[0])
}

func TestIsAtAll(t *testing.T) {
	must := assert.New(t)

	opts := new(options)
	var roll fw.Option
	roll = opts
	IsAtAll().Apply(roll)
	must.Equal(true, opts.isAtAll)
}

func TestMessageURL(t *testing.T) {
	must := assert.New(t)

	opts := new(options)
	var roll fw.Option
	roll = opts
	MessageURL("http://test.com").Apply(roll)
	must.Equal("http://test.com", opts.messageURL)
}

func TestPicURL(t *testing.T) {
	must := assert.New(t)

	opts := new(options)
	var roll fw.Option
	roll = opts
	PicURL("http://test.com/1.png").Apply(roll)
	must.Equal("http://test.com/1.png", opts.picURL)
}

func TestSingleTitle(t *testing.T) {
	must := assert.New(t)

	opts := new(options)
	var roll fw.Option
	roll = opts
	SingleTitle("test").Apply(roll)
	must.Equal("test", opts.singleTitle)
}

func TestSingleURL(t *testing.T) {
	must := assert.New(t)

	opts := new(options)
	var roll fw.Option
	roll = opts
	SingleURL("http://test.com").Apply(roll)
	must.Equal("http://test.com", opts.singleURL)
}

func TestBtnOrientation(t *testing.T) {
	must := assert.New(t)

	opts := new(options)
	var roll fw.Option
	roll = opts
	BtnOrientation("0").Apply(roll)
	must.Equal("0", opts.btnOrientation)
}

func TestHideAvatar(t *testing.T) {
	must := assert.New(t)

	opts := new(options)
	var roll fw.Option
	roll = opts
	HideAvatar("0").Apply(roll)
	must.Equal("0", opts.hideAvatar)
}
