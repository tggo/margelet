package margelet

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// MargeletSession incapsulates user's session
type margeletSession struct {
	*message
	bot         MargeletAPI
	lastMessage *tgbotapi.Message
	responses   []tgbotapi.Message
	finished    bool
}

func newMargetletSession(bot MargeletAPI, msg *tgbotapi.Message, responses []tgbotapi.Message) *margeletSession {
	return &margeletSession{
		message:   NewMessage(bot, msg),
		responses: responses,
		finished:  false,
	}
}

// Responses returns all user's responses in session
func (s *margeletSession) Responses() []tgbotapi.Message {
	return s.responses
}

func (s *margeletSession) Finish() {
	s.finished = true
}
