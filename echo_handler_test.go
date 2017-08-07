package margelet_test

import (
	"github.com/tggo/margelet"
)

// EchoHandler is simple handler example
type EchoHandler struct {
}

// Response send message back to author
func (handler EchoHandler) HandleMessage(m margelet.Message) error {
	_, err := m.QuickSend(m.Message().Text, nil)
	return err
}
