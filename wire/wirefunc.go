package wire

import (
	"github.com/google/wire"
)

func InitializeEvent(msg string) (Event, error) {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}, nil
}
