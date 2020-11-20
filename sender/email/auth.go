package email

import (
	"errors"
	"net/smtp"
)

type login struct {
	username string
	password string
}

// LoginAuth implements smtp.Auth which provides "LOGIN" mechanisms.
func LoginAuth(username, password string) smtp.Auth {
	return &login{username, password}
}

func (lg *login) Start(server *smtp.ServerInfo) (proto string, toServer []byte, err error) {
	return "LOGIN", []byte(lg.username), nil
}

func (lg *login) Next(fromServer []byte, more bool) (toServer []byte, err error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(lg.username), nil
		case "Password:":
			return []byte(lg.password), nil
		default:
			return nil, errors.New("Unknown fromServer")
		}
	}
	return
}
