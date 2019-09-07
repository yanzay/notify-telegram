package main

import (
	"fmt"
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
	env := []string{
		"GITHUB_REF", "GITHUB_SHA", "GITHUB_REPOSITORY",
		"GITHUB_ACTOR", "GITHUB_WORKFLOW", "GITHUB_HEAD_REF",
		"GITHUB_BASE_REF", "GITHUB_EVENT_NAME", "GITHUB_WORKSPACE",
		"GITHUB_ACTION", "GITHUB_EVENT_PATH", "RUNNER_OS",
		"RUNNER_TOOL_CACHE", "RUNNER_TEMP", "RUNNER_WORKSPACE",
	}
	for _, v := range env {
		fmt.Printf("%s: %s\n", v, os.Getenv(v))
	}
	_, err := c.SendMessage(chat, "test")
	if err != nil {
		log.Fatalf("unable to send message: %v", err)
	}
}
