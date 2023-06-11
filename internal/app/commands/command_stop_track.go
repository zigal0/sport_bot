package commands

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	StopTrackSuccessMsg = "You are now not tracked!"
)

func (svc *CommandService) StopTrack(inputMessage *tgbotapi.Message) {
	err := svc.commandProcessor.DeleteUser(inputMessage.From.ID)
	if err != nil {
		log.Printf("CommandService.StartTrack: error delete user from db: %v", err)
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, StopTrackSuccessMsg)

	_, err = svc.bot.Send(msg)
	if err != nil {
		log.Printf("CommandService.StopTrack: error sending reply message to chat: %v", err)
	}
}
