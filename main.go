package main

import (
	"net/http"
	"os"

	"github.com/yanzay/tbot/v2"
)

func main() {
	if len(os.Args) < 3 {
		println("too few arguments")
		return
	}
	var (
		token = os.Args[1]
		chat  = os.Args[2]
	)
	c := tbot.NewClient(token, http.DefaultClient, "https://api.telegram.org")
	_, err := c.SendMessage(chat, "test")
	if err != nil {
		println(err)
		return
	}
}
