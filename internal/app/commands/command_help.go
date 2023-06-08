package commands

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const helpMessage = `Available commands:
User commands:
/start_track - starts track user
/rename - changes username (default = telegram username)
/stop_track - deletes all user data and stops track

Exercise commands:
/create_exercise - creates new exercise with given name
/list_exercises - shows list of available exercises
/delete_exercise - deletes exercise by name

Activity commands:
/add_activity - adds your activity
/show_statistic_by_day - shows statistic by day 
/show_statistic_by_week - shows statistic by last week
/delete_activity - deletes activity by ID
`

func (svc *CommandService) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, helpMessage)

	_, err := svc.bot.Send(msg)
	if err != nil {
		log.Printf("CommandService.Help: error sending reply message to chat - %v", err)
	}
}
