package main

import (
	"flag"
	"log"
	"tg-bot/clients/telegram"
)

const (
	tgBorHost = "api.telegram.org"
)

func main() {
	tgClient := telegram.New(tgBorHost, mustToken())
}

func mustToken() string {
	token := flag.String("bot-token", "", "read bot token")
	flag.Parse()

	if *token == "" {
		log.Fatal("empty bot token")
	}

	return *token
}
