package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/Svoevolin/workshop_1_bot/internal/config"
	"github.com/Svoevolin/workshop_1_bot/internal/database"
	"github.com/Svoevolin/workshop_1_bot/internal/domain"
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

	if err = db.Migrator().DropTable(&domain.Rate{}, &domain.Expense{}, &domain.User{}); err != nil {
		log.Fatal("drop table failed", err)
	}
	if err = db.AutoMigrate(&domain.Rate{}, &domain.Expense{}, &domain.User{}); err != nil {
		log.Fatal("migrate failed", err)
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

	currencyExchangeRateWorker.Run(ctx)
	messageListenerWorker.Run(ctx)
}
