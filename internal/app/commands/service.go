package commands

import (
	"log"
	"runtime/debug"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/zigal0/sport_bot/internal/repository/users"
)

type CommandService struct {
	bot              *tgbotapi.BotAPI
	commandProcessor CommandProcessor

	usersRepo users.UsersRepo
}

func NewCommandService(
	bot *tgbotapi.BotAPI,
	usersRepo users.UsersRepo,
) *CommandService {
	return &CommandService{
		bot:       bot,
		usersRepo: usersRepo,
	}
}

func (svc *CommandService) HandleUpdate(update tgbotapi.Update) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("recovered from panic: %v\n%v", panicValue, string(debug.Stack()))
		}
	}()

	if update.Message != nil {
		svc.handleMessage(update.Message)
	}
}

func (svc *CommandService) handleMessage(inputMessage *tgbotapi.Message) {
	if !inputMessage.IsCommand() {
		svc.showCommandFormat(inputMessage)

		return
	}

	command := strings.Split(inputMessage.Command(), " ")[0]

	switch command {
	case "help":
		svc.Help(inputMessage)
	case "start_track":
		svc.StartTrack(inputMessage)
	}
}

func (svc *CommandService) showCommandFormat(inputMessage *tgbotapi.Message) {
	outputMsg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"Command format: /{command}")

	_, err := svc.bot.Send(outputMsg)
	if err != nil {
		log.Printf("CommandService.showCommandFormat: error sending reply message to chat - %v", err)
	}
}
