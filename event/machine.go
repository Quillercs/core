package event

import (
	"reflect"
	"runtime"
	"sync"
	"time"

	"github.com/quillercs/core/models"
	"github.com/quillercs/core/utils"
)

type eventFunc func(event models.Event)

type listener struct {
	sync.Mutex
	eventID   string
	eventFunc eventFunc
	once      bool
	done      bool
	UUID      string
}

type EventMachine struct {
	sync.Mutex
	queue     *utils.Queue
	listeners []*listener
}

func NewEventMachine() *EventMachine {
	return &EventMachine{
		queue: utils.NewQueue(),
	}
}

func (e *EventMachine) Push(event models.Event) {
	e.queue.Push(event)
}

func (e *EventMachine) Start() {
	for i := 0; i < 10; i++ {
		go func() {
			ticker := time.NewTicker(time.Second * 1)
			for now := range ticker.C {
				_ = now
				e.removeDone()
			}
		}()

		go func() {
			for {
				event := e.queue.Pop()
				if event != nil {
					go e.Dispatch(event.(models.Event))
				}

				runtime.Gosched()
				time.Sleep(time.Millisecond)
			}
		}()
	}
}

func (e *EventMachine) removeDone() {
	e.Lock()
	defer e.Unlock()

	i := 0
	for _, listener := range e.listeners {
		if listener != nil && listener.done {
			e.listeners, e.listeners[len(e.listeners)-1] = append(e.listeners[:i], e.listeners[i+1:]...), nil
		}
		i++
	}
}

func (e *EventMachine) Dispatch(event models.Event) {
	e.Lock()
	defer e.Unlock()

	for _, listener := range e.listeners {
		if listener.eventID == event.ID && !listener.done {
			listener.eventFunc(event)

			if listener.once {
				listener.done = true
			}
		}
	}
}

func (e *EventMachine) OnEvent(eventID string, eventFunc eventFunc) {
	e.Lock()
	defer e.Unlock()

	e.listeners = append(e.listeners, &listener{
		eventID:   eventID,
		eventFunc: eventFunc,
		once:      false,
	})
}

func (e *EventMachine) OnceEvent(eventID string, eventFunc eventFunc) {
	e.Lock()
	defer e.Unlock()

	e.listeners = append(e.listeners, &listener{
		eventID:   eventID,
		eventFunc: eventFunc,
		once:      true,
	})
}

func (e *EventMachine) Command(eventID, eventIDRes string, ef eventFunc) {
	ev := models.CreateEvent(eventID)
	e.queue.Push(ev)

	e.OnceEvent(eventIDRes, func(event models.Event) {
		if reflect.DeepEqual(event.UUID, ev.UUID) {
			ef(event)
		}
	})
}
