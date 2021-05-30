package controller

import "github.com/qwerty22121998/go-ascii/service"

type Provider struct {
	FromURL *URLToAsciiController
}

func NewProvider(provider *service.Provider) *Provider {
	return &Provider{
		FromURL: NewURLToAsciiController(provider),
	}
}
