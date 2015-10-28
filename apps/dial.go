package apps

import (
	appBase "github.com/quillercs/core/apps/base"
	"github.com/quillercs/core/event"
	"github.com/quillercs/core/models"
)

type Dial struct {
	appBase.App
}

func (d *Dial) Load(eventMachine *event.EventMachine) error {
	d.App.UID = "app::dial"
	d.App.Name = "Dial"
	d.App.MajorVersion = "1"
	d.App.MinorVersion = "0"
	d.App.Loaded = true
	d.EventMachine = eventMachine
	errChan := make(chan error)

	event.OnceEvent("event::core::loaded", func(event models.Event) {
	})

	// eventsEM.Listen(func(eventType EM.EventType, eventArgs EM.Args) {
	// 	if eventType == EM.EVENT_DOUBLE_TROUBLE {
	// 		errChan <- errors.New("DOUBLE TROUBLE, OH SHIT!")
	// 	} else if eventType == EM.EVENT_STATUS_OK {
	// 		errChan <- nil
	// 	}
	// })

	event.Push(models.NewEvent("get::core::status"))

	return <-errChan
}

func (d *Dial) Execute(args appBase.Args) error {
	err := EM.Dispatch(&EM.Args{
		Command: "event::app::dial",
		Args:    appBase.AppArgs,
	})

	return err
}

func (d *Dial) Unload() {
	d.Loaded = false
}
