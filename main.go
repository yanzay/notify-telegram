package main

import (
	"log"
	"net/http"
	"os"

	"github.com/yanzay/tbot/v2"
)

func main() {
	var (
		token = os.Getenv("INPUT_TOKEN")
		chat  = os.Getenv("INPUT_CHAT")
	)
	if token == "" {
		log.Fatal("token input is required")
	}
	if chat == "" {
		log.Fatal("chat input is required")
	}
	c := tbot.NewClient(token, http.DefaultClient, "https://api.telegram.org")
	_, err := c.SendMessage(chat, "test")
	if err != nil {
		log.Fatalf("unable to send message: %v", err)
	}
}
