package margelet

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type callbackQuery struct {
	*message
	query *tgbotapi.CallbackQuery
}

func newCallbackQuery(bot MargeletAPI, query *tgbotapi.CallbackQuery) *callbackQuery {
	return &callbackQuery{
		message: NewMessage(bot, query.Message),
		query:   query,
	}
}

// Query returns tgbotapi query
func (s *callbackQuery) Query() *tgbotapi.CallbackQuery {
	return s.query
}

func (s *callbackQuery) Data() string {
	return s.query.Data
}
