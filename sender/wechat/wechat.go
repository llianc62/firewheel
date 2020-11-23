// Package wechat provides sending message ability on WeChat.
package wechat

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"gopkg.in/resty.v1"

	fw "github.com/LiangXianSen/firewheel"
)

// WeChat implements Sender which sends message WeChat.
type WeChat struct {
	opts *options
}

// NewSender returns Sender which sends message to WeChat.
func NewSender(webhook, key string, opt ...fw.Option) fw.Sender {
	opts := new(options)
	opts.webhook = webhook
	opts.key = key

	var roll fw.Option
	roll = opts
	for _, o := range opt {
		roll = o.Apply(roll)
	}

	return &WeChat{
		opts: opts,
	}
}

// Send sends message to WeChat which you have to provides io.Reader include message body,
// and few options which implements Option.
func (wc *WeChat) Send(message io.Reader, opt ...fw.Option) (err error) {
	opts := new(options)
	*opts = *wc.opts

	var roll fw.Option
	roll = opts
	for _, o := range opt {
		roll = o.Apply(roll)
	}

	switch opts.message.MessageType {
	case "text":
		return wc.sendText(message, opts)
	case "markdown":
		return wc.sendMarkdown(message, opts)
	case "image":
		return wc.sendImage(message, opts)
	case "news":
		return wc.sendNews(message, opts)
	case "file":
		return wc.sendFile(message, opts)
	default:
		return wc.sendText(message, opts)
	}
}

func (wc *WeChat) sendText(message io.Reader, opts *options) (err error) {
	var content []byte
	if content, err = ioutil.ReadAll(message); err != nil {
		return
	}
	if opts.message.Text == nil {
		opts.message.Text = new(textMessage)
	}
	opts.message.Text.Content = string(content)

	return wc.sendMessage(opts.message)
}

func (wc *WeChat) sendMarkdown(message io.Reader, opts *options) (err error) {
	var content []byte
	if content, err = ioutil.ReadAll(message); err != nil {
		return
	}
	if opts.message.Markdown == nil {
		opts.message.Markdown = new(markdownMessage)
	}
	opts.message.Markdown.Content = string(content)

	return wc.sendMessage(opts.message)
}

func (wc *WeChat) sendNews(message io.Reader, opts *options) (err error) {
	return wc.sendMessage(opts.message)
}

func (wc *WeChat) sendImage(message io.Reader, opts *options) (err error) {
	var content []byte
	if content, err = ioutil.ReadAll(message); err != nil {
		return
	}
	if opts.message.Image == nil {
		opts.message.Image = new(imageMessage)
	}
	opts.message.Image.Base64 = base64.StdEncoding.EncodeToString(content)
	opts.message.Image.MD5Sum = fmt.Sprintf("%x", md5.Sum(content))

	return wc.sendMessage(opts.message)
}

func (wc *WeChat) sendFile(message io.Reader, opts *options) (err error) {
	var mediaID string
	if mediaID, err = wc.uploadFile(message, opts.filename); err != nil {
		return
	}
	if opts.message.File == nil {
		opts.message.File = new(fileMessage)
	}
	opts.message.File.MediaID = mediaID
	return wc.sendMessage(opts.message)
}

func (wc *WeChat) sendMessage(message Message) (err error) {
	var resp *resty.Response
	if resp, err = resty.SetTimeout(time.Second*3).R().
		SetQueryParam("key", wc.opts.key).
		SetHeader("Content-Type", "application/json").
		SetBody(message).
		Post(wc.opts.webhook); err != nil {
		return fmt.Errorf("http post failed: %s", err)
	}

	var response struct {
		ErrorCode    int    `json:"errcode"`
		ErrorMessage string `json:"errmsg"`
	}
	if err = json.Unmarshal(resp.Body(), &response); err != nil {
		return fmt.Errorf("parse response error: %s", err)
	}

	if resp.StatusCode() != http.StatusOK || response.ErrorCode != 0 {
		return fmt.Errorf("send message failed: %d, %s", response.ErrorCode, response.ErrorMessage)
	}
	return
}

func (wc *WeChat) uploadFile(fd io.Reader, filename string) (id string, err error) {
	var resp *resty.Response
	if resp, err = resty.SetTimeout(time.Second*3).R().
		SetQueryParam("key", wc.opts.key).
		SetQueryParam("type", "file").
		SetFileReader("media", filename, fd).
		Post("https://qyapi.weixin.qq.com/cgi-bin/webhook/upload_media"); err != nil {
		return
	}

	var media struct {
		ErrorCode    int    `json:"errcode"`
		ErrorMessage string `json:"errmsg"`
		MediaID      string `json:"media_id"`
	}
	if err = json.Unmarshal(resp.Body(), &media); err != nil {
		return "", fmt.Errorf("parse response error: %s", err)
	}

	if resp.StatusCode() != http.StatusOK || media.ErrorCode != 0 {
		return "", fmt.Errorf("upload failed: %d, %s", media.ErrorCode, media.ErrorMessage)
	}
	return media.MediaID, nil
}
