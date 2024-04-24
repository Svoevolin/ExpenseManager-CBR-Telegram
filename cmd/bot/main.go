package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/Svoevolin/workshop_1_bot/internal/config"
	"github.com/Svoevolin/workshop_1_bot/internal/database"
	"github.com/Svoevolin/workshop_1_bot/internal/infrastructure/cbr_gateway"
	"github.com/Svoevolin/workshop_1_bot/internal/infrastructure/tg_gateway"
	"github.com/Svoevolin/workshop_1_bot/internal/model/messages"
	"github.com/Svoevolin/workshop_1_bot/internal/services"
	"github.com/Svoevolin/workshop_1_bot/internal/worker"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	config, err := config.New()
	if err != nil {
		log.Fatal("config init failed:", err)
	}

	// DATABASE

	db, err := gorm.Open(postgres.Open("host=localhost port=5432 user=postgres password=postgres"))
	if err != nil {
		log.Fatal("database init failed", err)
	}

	userDB := database.NewUserDB(db)

	rateDB := database.NewRateDB(db)

	expenseDB := database.NewExpenseDB(db)

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

	var wg sync.WaitGroup
	wg.Add(2)

	go currencyExchangeRateWorker.Run(ctx, &wg)
	go messageListenerWorker.Run(ctx, &wg)

	wg.Wait()
}
