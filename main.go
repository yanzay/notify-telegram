package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/yanzay/tbot/v2"
)

var icons = map[string]string{
	"failure":   "❗",
	"cancelled": "❕",
	"success":   "✅",
}

func main() {
	var (
		token  = os.Getenv("INPUT_TOKEN")
		chat   = os.Getenv("INPUT_CHAT")
		status = os.Getenv("INPUT_STATUS")
	)
	if token == "" {
		log.Fatal("token input is required")
	}
	if chat == "" {
		log.Fatal("chat input is required")
	}
	if status == "" {
		log.Fatal("status input is required")
	}
	c := tbot.NewClient(token, http.DefaultClient, "https://api.telegram.org")

	icon := icons[strings.ToLower(status)]

	workflow := os.Getenv("GITHUB_WORKFLOW")
	repo := os.Getenv("GITHUB_REPOSITORY")
	msg := fmt.Sprintf(`%s*%s*: %s (%s)`, icon, status, repo, workflow)

	_, err := c.SendMessage(chat, msg, tbot.OptParseModeMarkdown)
	if err != nil {
		log.Fatalf("unable to send message: %v", err)
	}
}
