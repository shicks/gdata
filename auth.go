package gdata

import (
	"os"
	"url"
)

type Auth interface {
	Authenticate(*Request) os.Error
}

type authSub struct {
	token string
}

type oAuth struct {
	token, secret string
}

func (a authSub) Authenticate(req *Request) os.Error {
	
}

func NewOAuth() Auth {
	return OAuth{
		user: user,
		pass: password,
	}
}

type CaptchaHandler func(string) string

func NewClientLogin(user, password string, handleCaptcha CaptchaHandler) Auth {
	return authSub{token: ""}
}