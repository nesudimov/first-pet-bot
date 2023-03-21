package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/nesudimov/first-pet-bot/clients/telegram"
)

const (
	tgBotHost = "api.telegram.org"
	token     = "6219465977:AAGAaZsk_sQGphPz2xRXjSv6jdrQJLhCHkI"
)

func main() {

	tgClient := telegram.New(tgBotHost, mustToken())

	// fetcher = fetcher.New()

	// processor = processor.New()

	// cosumer.Start(fetcher, processor)

	tgApiUrl := "https://api.telegram.org/bot"
	method := "/getUpdates"
	url := tgApiUrl + token + method

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	fmt.Println(string(body))
}

func mustToken() string {
	token := flag.String(
		"token-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified...")
	}

	return *token
}
