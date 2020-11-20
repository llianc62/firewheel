package dingding_test

/*
import (
	"fmt"
	"strings"
	"testing"

	"github.com/royeo/dingrobot"
	"github.com/stretchr/testify/assert"

	fw "github.com/LiangXianSen/firewheel"
	"github.com/LiangXianSen/firewheel/sender/dingding"
)

func TestDingDingSend(t *testing.T) {
	var sender fw.Sender

	must := assert.New(t)
	webhook := "https://oapi.dingtalk.com/robot/send?access_token="
	sender = dingding.NewSender(webhook)
	err := sender.Send(strings.NewReader("keywords: test message"))
	must.Nil(err)
}

func TestDingDingSendMarkdown(t *testing.T) {
	var sender fw.Sender

	must := assert.New(t)
	keywords := "keywords"
	webhook := "https://oapi.dingtalk.com/robot/send?access_token="
	sender = dingding.NewSender(webhook)
	pattern := `
	## topic: license expired
	something happened, pls checkout.
	## detail:
	- one
	- two
	- three
	`

	err := sender.Send(
		strings.NewReader(pattern),
		dingding.MessageMode(dingding.MarkdownMode),
		dingding.KeyWords(keywords),
	)
	must.Nil(err)
}

func TestDingDingSendLink(t *testing.T) {
	var sender fw.Sender

	must := assert.New(t)
	keywords := "keywords"
	webhook := "https://oapi.dingtalk.com/robot/send?access_token="
	sender = dingding.NewSender(webhook)
	err := sender.Send(
		strings.NewReader("test message"),
		dingding.MessageMode(dingding.LinkMode),
		dingding.KeyWords(keywords),
		dingding.MessageURL("https://gist.github.com/asim/d3dea1832609b32538838529e693f9bf"),
		dingding.PicURL("https://ss0.bdstatic.com/70cFvHSh_Q1YnxGkpoWK1HF6hhy/it/u=110199744,2712833928&fm=26&gp=0.jpg"),
	)
	must.Nil(err)
}

func TestDingDingSendActionCard(t *testing.T) {
	var sender fw.Sender

	must := assert.New(t)
	keywords := "keywords"
	webhook := "https://oapi.dingtalk.com/robot/send?access_token="
	sender = dingding.NewSender(webhook)
	err := sender.Send(
		strings.NewReader("test message"),
		dingding.MessageMode(dingding.ActionCardMode),
		dingding.KeyWords(keywords),
		dingding.SingleTitle("read more"),
		dingding.SingleURL("https://gist.github.com/asim/d3dea1832609b32538838529e693f9bf"),
	)
	must.Nil(err)
}

func TestDingRobotWithText(t *testing.T) {
	must := assert.New(t)

	var atMobiles []string
	var isAtAll bool

	// Content must contains keywords
	keywords := "keywords"
	webhook := "https://oapi.dingtalk.com/robot/send?access_token="
	robot := dingrobot.NewRobot(webhook)

	content := fmt.Sprintf("%s: test message", keywords)
	err := robot.SendText(content, atMobiles, isAtAll)
	must.Nil(err)
}

func TestDingRobotWithMarkdown(t *testing.T) {
	must := assert.New(t)

	var atMobiles []string
	var isAtAll bool

	// Content must contains keywords
	keywords := "keywords"
	webhook := "https://oapi.dingtalk.com/robot/send?access_token="
	robot := dingrobot.NewRobot(webhook)

	pattern := `
	## topic: license expired
	something happened, pls checkout.
	## detail:
	- one
	- two
	- three
	`

	err := robot.SendMarkdown(keywords, pattern, atMobiles, isAtAll)
	must.Nil(err)
}

func TestDingRobotWithLink(t *testing.T) {
	must := assert.New(t)

	// Content must contains keywords
	keywords := "keywords"
	webhook := "https://oapi.dingtalk.com/robot/send?access_token="
	robot := dingrobot.NewRobot(webhook)
	content := "test message"
	messageURL := "https://gist.github.com/asim/d3dea1832609b32538838529e693f9bf"
	picURL := "https://ss0.bdstatic.com/70cFvHSh_Q1YnxGkpoWK1HF6hhy/it/u=110199744,2712833928&fm=26&gp=0.jpg"
	err := robot.SendLink(keywords, content, messageURL, picURL)
	must.Nil(err)
}

func TestDingRobotWithActionCard(t *testing.T) {
	must := assert.New(t)

	// Content must contains keywords
	keywords := "keywords"
	webhook := "https://oapi.dingtalk.com/robot/send?access_token="
	robot := dingrobot.NewRobot(webhook)
	content := "test message"
	singleTitle := "read more"
	singleURL := "https://gist.github.com/asim/d3dea1832609b32538838529e693f9bf"
	btnOrientation := "0"
	hideAvatar := "0"
	err := robot.SendActionCard(keywords, content, singleTitle, singleURL, btnOrientation, hideAvatar)
	must.Nil(err)
}
*/
