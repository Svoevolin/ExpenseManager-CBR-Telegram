package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/Svoevolin/workshop_1_bot/internal/config"
	"github.com/Svoevolin/workshop_1_bot/internal/database"
	"github.com/Svoevolin/workshop_1_bot/internal/infrastructure/cbr_gateway"
	"github.com/Svoevolin/workshop_1_bot/internal/infrastructure/tg_gateway"
	"github.com/Svoevolin/workshop_1_bot/internal/model/messages"
	"github.com/Svoevolin/workshop_1_bot/internal/services"
	"github.com/Svoevolin/workshop_1_bot/internal/worker"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	config, err := config.New()
	if err != nil {
		log.Fatal("config init failed:", err)
	}

	// DATABASE

	userDB, err := database.NewUserDB()
	if err != nil {
		log.Fatal("database init failed", err)
	}

	rateDB, err := database.NewRateDB()
	if err != nil {
		log.Fatal("database init failed", err)
	}

	expenseDB, err := database.NewExpenseDB()
	if err != nil {
		log.Fatal("database init failed", err)
	}

	// GATEWAY

	cbrRateAPIGateway := cbr_gateway.New()
	tgAPIgateaway, err := tg_gateway.New(config)
	if err != nil {
		log.Fatal("tg client init failed:", err)
	}

	// SERVICES

	exchangeRateUpdateSvc := services.NewExchangeRateUpdateSvc(cbrRateAPIGateway, rateDB, config)

	// COMMANDS

	messageProcessor := messages.New(tgAPIgateaway, config, userDB, rateDB, expenseDB, exchangeRateUpdateSvc)

	// WORKERS

	currencyExchangeRateWorker := worker.NewCurrencyExchangeRateWorker(exchangeRateUpdateSvc, config)
	messageListenerWorker := worker.NewMessageListenerWorker(tgAPIgateaway, messageProcessor)

	currencyExchangeRateWorker.Run(ctx)
	messageListenerWorker.Run(ctx)
}
