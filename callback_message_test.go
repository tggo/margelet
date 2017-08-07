package margelet_test

import (
	"github.com/tggo/margelet"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type CallbackMessage struct {
}

func (handler CallbackMessage) HandleCallback(query margelet.CallbackQuery) error {
	config := tgbotapi.CallbackConfig{
		CallbackQueryID: query.Query().ID,
		Text:            "Done!",
		ShowAlert:       false,
	}

	query.Bot().AnswerCallbackQuery(config)
	return nil
}
