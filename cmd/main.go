package main

import (
	"github.com/qwerty22121998/go-ascii/pkg/server"
)

func main() {
	s := server.NewServer()
	s.Register()
	s.Start()
}
