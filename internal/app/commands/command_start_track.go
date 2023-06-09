package commands

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/zigal0/sport_bot/internal/domain"
)

const (
	StartTrackSuccessMsg = "You are now tracked!"
)

func (svc *CommandService) StartTrack(inputMessage *tgbotapi.Message) {
	user := domain.User{
		ID:       inputMessage.From.ID,
		Username: inputMessage.From.UserName,
	}
	
	err := svc.usersRepo.AddUser(user)
	if err != nil {
		log.Print("CommandService.StartTrack: error add user to db")
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, StartTrackSuccessMsg)

	_, err = svc.bot.Send(msg)
	if err != nil {
		log.Printf("CommandService.Help: error sending reply message to chat - %v", err)
	}
}
