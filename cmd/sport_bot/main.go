package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/zigal0/sport_bot/internal/app/commands"
	"github.com/zigal0/sport_bot/internal/repository/users"
)

const (
	botTokenKey = "BOT_TOKEN"
	pathToENV   = "env/.env"
)

func main() {
	err := godotenv.Load(pathToENV)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token, found := os.LookupEnv(botTokenKey)
	if !found {
		log.Panic("Environment variable TOKEN not found in .env")
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

	sportRepo := users.NewUsersRepo()

	commandSvc := commands.NewCommandService(bot, sportRepo)

	for update := range updates {
		commandSvc.HandleUpdate(update)
	}
}
