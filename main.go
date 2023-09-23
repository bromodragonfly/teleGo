package main

import (
	"flag"
	"log"
	"telego/clients/telegram"
)

//TODO: get HOST from flag
const (
	tgBotHost = "api.telegram.org"
)

func main() {
	tgClient := telegram.New(tgBotHost, mustToken())

}

func mustToken() string {
	token := flag.String("token-bot-token", "", "token for acces to TG bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
