package event

type listeners []*listener

type eventFunc func(event Event)

type listener struct {
	eventFunc
	eventID string
	once    bool
	done    bool
}
