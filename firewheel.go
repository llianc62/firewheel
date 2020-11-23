// Package firewheel defines a message delivery model, provides various platform
// warpped sdk, such as Mail, DingTalk, WeChat and etc.
package firewheel

import (
	"io"
)

// Messager binds sender and indicates where the data source come from.
type Messager interface {
	// Sender requires send ability.
	Sender

	// Setup binds sender.
	Setup(...Sender)
}

// Sender provides message delivery ability.
type Sender interface {
	// Send sends message which is the main processor
	// requires io.Reader include message body and
	// several args which implements Option.
	Send(io.Reader, ...Option) error
}
