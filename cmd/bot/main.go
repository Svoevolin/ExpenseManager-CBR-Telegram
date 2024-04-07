package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/Svoevolin/workshop_1_bot/internal/config"
	"github.com/Svoevolin/workshop_1_bot/internal/database"
	"github.com/Svoevolin/workshop_1_bot/internal/infrastructure/tg_gateaway"
	"github.com/Svoevolin/workshop_1_bot/internal/model/messages"
	"github.com/Svoevolin/workshop_1_bot/internal/worker"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	config, err := config.New()
	if err != nil {
		log.Fatal("config init failed:", err)
	}

	tgAPIgateaway, err := tg_gateaway.New(config)
	if err != nil {
		log.Fatal("tg client init failed:", err)
	}

	userDB, err := database.NewUserDB()
	if err != nil {
		log.Fatal("database init failed", err)
	}

	messageProcessor := messages.New(tgAPIgateaway, config, userDB)
	messageListenerWorker := worker.NewMessageListenerWorker(tgAPIgateaway, messageProcessor)
	messageListenerWorker.Run(ctx)
}
