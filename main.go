package main

import (
	"geek/login"
)

func main() {
	c := login.NewLoginClient()

	result := c.Login("1", "1")
	if !result.IsLoginSuccess() {
		//return "", "", "", errors.New(result.Error.Msg)
	}
}
