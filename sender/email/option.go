package email

import (
	"crypto/tls"
	"io"
	"time"

	fw "github.com/LiangXianSen/firewheel"
)

// options implements Option.
type options struct {
	poolSize    int
	timeout     time.Duration
	tlsConf     []*tls.Config
	contentType string
	from        string
	to          []string
	bcc         []string
	cc          []string
	subject     string
	attach      *attachment

	f func(*options)
}

func (opts *options) Apply(opt fw.Option) fw.Option {
	opts.f(opt.(*options))
	return opt
}

// PoolSize sets pool size.
func PoolSize(size int) fw.Option {
	return &options{
		f: func(o *options) {
			o.poolSize = size
		},
	}
}

// Timeout sets timeout.
func Timeout(timeout time.Duration) fw.Option {
	return &options{
		f: func(o *options) {
			o.timeout = timeout
		},
	}
}

// ServerName sets TLS Config of Server Name.
func ServerName(name string) fw.Option {
	return &options{
		f: func(o *options) {
			o.tlsConf = append(o.tlsConf, &tls.Config{ServerName: name})
		},
	}
}

// HTMLContentType sets HTML message type.
func HTMLContentType() fw.Option {
	return &options{
		f: func(o *options) {
			o.contentType = "HTML"
		},
	}
}

// TextContentType sets Text message type.
func TextContentType() fw.Option {
	return &options{
		f: func(o *options) {
			o.contentType = "Text"
		},
	}
}

// Subject sets email subject.
func Subject(subject string) fw.Option {
	return &options{
		f: func(o *options) {
			o.subject = subject
		},
	}
}

// SubjectFunc sets email subject from the func returned.
func SubjectFunc(fn func() string) fw.Option {
	return &options{
		f: func(o *options) {
			o.subject = fn()
		},
	}
}

// From sets email from address.
func From(from string) fw.Option {
	return &options{
		f: func(o *options) {
			o.from = from
		},
	}
}

// To sets email To addresses.
func To(to ...string) fw.Option {
	return &options{
		f: func(o *options) {
			o.to = to
		},
	}
}

// Cc sets email Cc addresses.
func Cc(cc ...string) fw.Option {
	return &options{
		f: func(o *options) {
			o.cc = cc
		},
	}
}

// Bcc sets email Bcc addresses.
func Bcc(bcc ...string) fw.Option {
	return &options{
		f: func(o *options) {
			o.bcc = bcc
		},
	}
}

type attachment struct {
	filename string
	fd       io.Reader
}

// Attachment sets email attachment.
func Attachment(fd io.Reader, filename string) fw.Option {
	return &options{
		f: func(o *options) {
			o.attach = &attachment{
				filename: filename,
				fd:       fd,
			}
		},
	}
}
