package event

import (
	"time"

	"github.com/satori/go.uuid"
)

type Event struct {
	ID   string `json:"ID"`
	Date int64  `json:"date"`
	Data []byte `json:"data"`
	UUID []byte `json:"uuid"`
}

func CreateEvent(id string) Event {
	now := time.Now()
	return Event{
		ID: id, Date: now.UnixNano(),
		Data: []byte(""), UUID: uuid.NewV4().Bytes(),
	}
}

func CreateResponseEvent(otherEvent Event) Event {
	now := time.Now()
	return Event{
		ID: otherEvent.ID + "::response", Date: now.UnixNano(),
		Data: []byte(""), UUID: otherEvent.UUID,
	}
}

func CreateEventWithData(id string, data []byte) Event {
	now := time.Now()
	return Event{
		ID: id, Date: now.UnixNano(),
		Data: data, UUID: uuid.NewV4().Bytes(),
	}
}
