package main

import (
	"flag"
	"github.com/tumbleweedd/mySecondTestVKbot/pkg/telegram"
	"log"
)

//const botToken = "6279047527:AAEtZGFDHKToNXk4tORbWqft-f0Krm-Jswc"

func mustToken() string {
	token := flag.String(
		"bot-token",
		"",
		"token for access",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}

func main() {
	token := mustToken()
	botApi := "https://api.telegram.org/bot"
	botUrl := botApi + token

	telegram.Run(botUrl)
}
