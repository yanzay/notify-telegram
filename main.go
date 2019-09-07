package main

import (
	"log"
	"net/http"
	"os"

	"github.com/yanzay/tbot/v2"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("too few arguments")
	}
	var (
		token = os.Args[1]
		chat  = os.Args[2]
	)
	c := tbot.NewClient(token, http.DefaultClient, "https://api.telegram.org")
	_, err := c.SendMessage(chat, "test")
	if err != nil {
		log.Fatalf("unable to send message: %v", err)
	}
}
