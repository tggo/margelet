package margelet_test

import (
	"net/url"

	"github.com/tggo/margelet"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type BotMock struct {
	Updates chan tgbotapi.Update
}

func (bot BotMock) Send(c tgbotapi.Chattable) (tgbotapi.Message, error) {
	return tgbotapi.Message{}, nil
}

func (bot BotMock) AnswerInlineQuery(config tgbotapi.InlineConfig) (tgbotapi.APIResponse, error) {
	return tgbotapi.APIResponse{}, nil
}

func (bot BotMock) AnswerCallbackQuery(config tgbotapi.CallbackConfig) (tgbotapi.APIResponse, error) {
	return tgbotapi.APIResponse{}, nil
}

func (bot BotMock) GetFileDirectURL(fileID string) (string, error) {
	return "https://example.com/test.txt", nil
}

func (bot BotMock) IsMessageToMe(message tgbotapi.Message) bool {
	return false
}

func (bot BotMock) GetUpdatesChan(config tgbotapi.UpdateConfig) (tgbotapi.UpdatesChannel, error) {
	return bot.Updates, nil
}

func (bot BotMock) MakeRequest(endpoint string, params url.Values) (tgbotapi.APIResponse, error) {
	return tgbotapi.APIResponse{}, nil
}

var (
	botMock = BotMock{}
)

func getMargelet() *margelet.Margelet {
	botMock.Updates = make(chan tgbotapi.Update, 10)
	m, _ := margelet.NewMargeletFromBot("test", "127.0.0.1:6379", "", 10, &botMock, false)

	m.Redis.FlushDb()
	return m
}

//Empty Function, because "go vet" wants Margelet() to exist because the ExampleMargelet function
func Margelet() {
}

func ExampleMargelet() {
	bot, err := margelet.NewMargelet("<your awesome bot name>", "<redis addr>", "<redis password>", 0, "your bot token", false)

	if err != nil {
		panic(err)
	}

	bot.Run()
}
