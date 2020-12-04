# firewheel 
[![Build Status](https://travis-ci.org/LiangXianSen/firewheel.svg?branch=main)](https://travis-ci.org/LiangXianSen/firewheel)
[![codecov](https://codecov.io/gh/LiangXianSen/firewheel/branch/main/graph/badge.svg?token=HK8JOVPZN3)](https://codecov.io/gh/LiangXianSen/firewheel)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/LiangXianSen/firewheel)](https://pkg.go.dev/github.com/LiangXianSen/firewheel)
[![CodeFactor](https://www.codefactor.io/repository/github/liangxiansen/firewheel/badge)](https://www.codefactor.io/repository/github/liangxiansen/firewheel)
![GitHub](https://img.shields.io/github/license/LiangXianSen/firewheel)

Package firewheel defines a message delivery model, provides various platform warpped sdk, such as Email, DingTalk, WeChat and etc.



Features:

- [x] Email
- [x] DingTalk
- [x] WeChat



# Installation

```
go get -u github.com/LiangXianSen/firewheel
```



# Guide

firewheel is well designed as out-of-the-box, you can send message by `Sender`  which you choose, or implements a `Messager`, sets up multiple senders. 



Simply usage: 

```go
package main

import (
	"log"
	"strings"

	fw "github.com/LiangXianSen/firewheel"
	"github.com/LiangXianSen/firewheel/sender/email"
)

func main() {
    var err error
    var sender fw.Sender

    if sender, err = email.NewSender(
      "smtp.gmail.com:587",                     // smtp server
      email.LoginAuth("username", "password"),  // smtp.Auth 
    ); err != nil {
      log.Fatal(err)
    }

    if err = sender.Send(
      strings.NewReader("test message"),        // io.Reader include message body
      email.TextContentType(),
      email.Subject("email test"),
      email.From("test@gmail.com"),
      email.To("test@gmail.com", "example@gmail.com"),
      email.Bcc("example@gmail.com"),
      email.Cc("example@gmail.com"),
    ); err != nil {
      log.Fatal(err)
    }
}
```

there are some tips: 

- You can gives all options when new  a sender, use each  `Send()` function will include all options once you provided. On condition that you don't usually change.

  ```go
  sender, err = email.NewSender(
  		"smtp.gmail.com:587",
  		email.LoginAuth("username", "password"),
  		email.TextContentType(),
  		email.Subject("email test"),
  		email.From("test@gmail.com"),
  		email.To("example@gmail.com"),
  		email.Bcc("example@gmail.com"),
  		email.Cc("example@gmail.com"),
  	)
  ```

  or at some conditions change some of them.

  ```go
  sender.Send(
  	strings.NewReader("test message"),
  	email.Subjectc("new email test"),
  )
  ```

  

- sends message with dynamic subject:

  ```go
  sender, err = email.NewSender(
  	        "smtp.gmail.com:587",
  	        email.LoginAuth("username", "password"),
  	        email.TextContentType(),
  	        email.SubjectFunc(func() string {
  	        	date := time.Now().Format("2016-01-02 15:04:05")
  	        	return "test date: " + date
  	        })
  	)
  ```





Massager pattern:

You also can implements `Massager` , sets up few senders. Sends one mssage to different destination. [see examples](https://github.com/LiangXianSen/firewheel/tree/main/examples)



