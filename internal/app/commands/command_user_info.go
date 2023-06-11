package commands

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (svc *CommandService) UserInfo(inputMessage *tgbotapi.Message) {
	user, err := svc.commandProcessor.GetUser(inputMessage.From.ID)
	if err != nil {
		log.Printf("CommandService.UserInfo: error get user from db: %v", err)
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("You are %s. Get faster, higher, stronger!", user.Username),
	)

	_, err = svc.bot.Send(msg)
	if err != nil {
		log.Printf("CommandService.UserInfo: error sending reply message to chat: %v", err)
	}
}
