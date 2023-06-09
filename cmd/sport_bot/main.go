package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/zigal0/sport_bot/internal/app/commands"
	"github.com/zigal0/sport_bot/internal/repository/users"
)

const (
	botTokenKey = "BOT_TOKEN"
	pathToENV   = "env/.env"
	pgDSNKey    = "PG_DSN"
)

func main() {
	err := godotenv.Load(pathToENV)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	pgDSN, found := os.LookupEnv(pgDSNKey)
	if !found {
		log.Panic("Environment variable PG_DSN not found in .env")
	}

	db, err := sqlx.Connect("pgx", pgDSN)
	if err != nil {
		log.Fatalln(err)
	}

	token, found := os.LookupEnv(botTokenKey)
	if !found {
		log.Panic("Environment variable BOT_TOKEN not found in .env")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}

	// Uncomment if you want debugging
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	updateCfg := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates := bot.GetUpdatesChan(updateCfg)

	sportRepo := users.NewUsersRepo(db)

	commandSvc := commands.NewCommandService(bot, sportRepo)

	for update := range updates {
		commandSvc.HandleUpdate(update)
	}
}
