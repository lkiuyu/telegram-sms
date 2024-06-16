package handler

import (
	"fmt"

	"github.com/damonto/telegram-sms/internal/pkg/util"
	"gopkg.in/telebot.v3"
)

type StartHandler struct{}

func HandleStartCommand(c telebot.Context) error {
	message := `
Hello, *%s %s*!
Thanks for using this bot.
Your UID is *%d*
	`
	return c.Send(util.EscapeText(fmt.Sprintf(message, c.Sender().FirstName, c.Sender().LastName, c.Sender().ID)), &telebot.SendOptions{
		ParseMode: telebot.ModeMarkdownV2,
	})
}
