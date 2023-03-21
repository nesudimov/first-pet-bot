package main

import (
	"flag"
	"log"

	tg "github.com/nesudimov/first-pet-bot/clients/telegram"
	event_consumer "github.com/nesudimov/first-pet-bot/consumer/event-consumer"
	"github.com/nesudimov/first-pet-bot/events/telegram"
	"github.com/nesudimov/first-pet-bot/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

func main() {

	eventsProcessor := telegram.New(
		tg.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)
	log.Printf("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("Service is stopped", err)
	}
}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified...")
	}

	return *token
}
