package email

import (
	"strings"
	"testing"
	"time"

	fw "github.com/LiangXianSen/firewheel"
	"github.com/stretchr/testify/assert"
)

func TestPoolSize(t *testing.T) {
	must := assert.New(t)

	opts := new(options)
	var roll fw.Option
	roll = opts
	PoolSize(4).Apply(roll)
	must.Equal(4, opts.poolSize)
}

func TestTimeout(t *testing.T) {
	must := assert.New(t)

	opts := new(options)
	var roll fw.Option
	roll = opts
	Timeout(10 * time.Second).Apply(roll)
	must.Equal(10*time.Second, opts.timeout)
}

func TestServerName(t *testing.T) {
	must := assert.New(t)

	opts := new(options)
	var roll fw.Option
	roll = opts
	ServerName("smtp.gmail.com").Apply(roll)
	must.Equal("smtp.gmail.com", opts.tlsConf[0].ServerName)
}

func TestHTMLContentType(t *testing.T) {
	must := assert.New(t)

	opts := new(options)
	var roll fw.Option
	roll = opts
	HTMLContentType().Apply(roll)
	must.Equal("HTML", opts.contentType)
}

func TestTextContentType(t *testing.T) {
	must := assert.New(t)

	opts := new(options)
	var roll fw.Option
	roll = opts
	TextContentType().Apply(roll)
	must.Equal("Text", opts.contentType)
}

func TestSubject(t *testing.T) {
	must := assert.New(t)

	opts := new(options)
	var roll fw.Option
	roll = opts
	Subject("test").Apply(roll)
	must.Equal("test", opts.subject)
}

func TestSubjectFunc(t *testing.T) {
	must := assert.New(t)

	opts := new(options)
	var roll fw.Option
	roll = opts
	SubjectFunc(func() string {
		return "test-" + time.Now().Format(time.RFC3339)
	}).Apply(roll)
	must.NotEmpty(opts.subject)
}

func TestFrom(t *testing.T) {
	must := assert.New(t)

	opts := new(options)
	var roll fw.Option
	roll = opts
	From("test@gmail.com").Apply(roll)
	must.Equal("test@gmail.com", opts.from)
}

func TestTo(t *testing.T) {
	must := assert.New(t)

	opts := new(options)
	var roll fw.Option
	roll = opts
	To("example@gmail.com").Apply(roll)
	must.Equal("example@gmail.com", opts.to[0])
}

func TestCc(t *testing.T) {
	must := assert.New(t)

	opts := new(options)
	var roll fw.Option
	roll = opts
	Cc("example@gmail.com").Apply(roll)
	must.Equal("example@gmail.com", opts.cc[0])
}

func TestBcc(t *testing.T) {
	must := assert.New(t)

	opts := new(options)
	var roll fw.Option
	roll = opts
	Bcc("example@gmail.com").Apply(roll)
	must.Equal("example@gmail.com", opts.bcc[0])
}

func TesAttachment(t *testing.T) {
	must := assert.New(t)

	opts := new(options)
	var roll fw.Option
	roll = opts
	Attachment(strings.NewReader("test"), "test.txt").Apply(roll)
	must.NotNil(opts.attachment)
}
