package event

import (
	"log"
	"reflect"
	"runtime"
	"sync"
	"time"
)

type Machine struct {
	sync.RWMutex
	events chan Event
	listeners
}

func NewMachine() *Machine {
	m := &Machine{}
	m.events = make(chan Event)
	return m
}

func (m *Machine) Start() {
	go func() {
		ticker := time.NewTicker(time.Second * 1)
		for now := range ticker.C {
			_ = now
			m.removeDone()
		}
	}()

	go func() {
		for {
			select {
			case event := <-m.events:
				go m.dispatch(event)
			default:
				runtime.Gosched()
				time.Sleep(time.Millisecond)
			}
		}
	}()

	m.Push(CreateEvent("core::running"))

	m.OnceEvent("core::status", func(ev Event) {
		log.Println("Response status", ev)
		newEvent := CreateResponseEvent(ev)
		newEvent.Data = []byte("Running")
		m.Push(newEvent)
	})
}

func (m *Machine) dispatch(event Event) {
	m.Lock()
	defer m.Unlock()

	for _, listener := range m.listeners {
		if listener.eventID == event.ID && !listener.done {
			listener.eventFunc(event)

			if listener.once {
				listener.done = true
			}
		}
	}
}

func (m *Machine) removeDone() {
	m.Lock()
	defer m.Unlock()

	i := 0
	for _, listener := range m.listeners {
		if listener != nil && listener.done {
			m.listeners, m.listeners[len(m.listeners)-1] = append(m.listeners[:i], m.listeners[i+1:]...), nil
		}
		i++
	}
}

func (m *Machine) Push(event Event) {
	go func() {
		m.events <- event
	}()
}

func (m *Machine) OnceEvent(eventID string, eventFunc eventFunc) {
	m.Lock()
	defer m.Unlock()

	m.listeners = append(m.listeners, &listener{
		eventID:   eventID,
		eventFunc: eventFunc,
		once:      true,
	})
}

func (m *Machine) OnEvent(eventID string, eventFunc eventFunc) {
	m.Lock()
	defer m.Unlock()

	m.listeners = append(m.listeners, &listener{
		eventID:   eventID,
		eventFunc: eventFunc,
	})
}

func (m *Machine) Command(eventID string, eventFunc eventFunc) {
	ev := CreateEvent(eventID)

	m.OnceEvent(eventID+"::response", func(event Event) {
		if reflect.DeepEqual(event.UUID, ev.UUID) {
			eventFunc(event)
		}
	})

	m.Push(ev)
}
