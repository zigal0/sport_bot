package commands

import (
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/zigal0/sport_bot/internal/domain"
)

const (
	MsgNeedMoreArgsForRename = "You need to enter new username with /rename_user command"
)

func (svc *CommandService) RenameUser(inputMessage *tgbotapi.Message) {
	commandContent := strings.Split(inputMessage.Text, " ")
	if len(commandContent) < 2 {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, MsgNeedMoreArgsForRename)

		_, err := svc.bot.Send(msg)
		if err != nil {
			log.Printf("CommandService.RenameUser: error sending reply message to chat: %v", err)
		}

		return
	}

	newUserName := strings.Join(commandContent[1:], " ")

	user := domain.User{
		ID:       inputMessage.From.ID,
		Username: newUserName,
	}

	err := svc.commandProcessor.UpdateUser(user)
	if err != nil {
		log.Printf("CommandService.RenameUser: error update user in db: %v", err)
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("Your username now is %s", newUserName),
	)

	_, err = svc.bot.Send(msg)
	if err != nil {
		log.Printf("CommandService.RenameUser: error sending reply message to chat: %v", err)
	}
}
