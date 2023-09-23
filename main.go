package main

import (
	"flag"
	"log"
)

func main() {
	t, err := mustToken()
}

func mustToken() string {
	token := flag.String("token-bot-token", "", "token for acces to TG bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
