package generic

import (
	// EM "github.com/quillercs/eventmachine"
	"errors"

	appBase "github.com/quillercs/core/apps/base"
)

type Dial struct {
	appBase.App
}

func (d *Dial) Load(args appBase.Args) error {
	d.App.UID = "app::dial"
	d.App.Name = "Dial"
	d.App.MajorVersion = "1"
	d.App.MinorVersion = "0"
	d.App.Loaded = true

	errChan := make(chan error)

	EM.Listen(func(eventType EM.EventType, eventArgs EM.Args) {
		if eventType == EM.EVENT_DOUBLE_TROUBLE {
			errChan <- errors.New("DOUBLE TROUBLE, OH SHIT!")
		} else if eventType == EM.EVENT_STATUS_OK {
			errChan <- nil
		}
	})

	EM.Dispatch(&EM.Args{
		Command: "event::core::status",
	})

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
