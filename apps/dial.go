package apps

type Dial struct {
	App
}

func (d *Dial) Load() error {
	d.App.UID = "app::dial"
	d.App.Name = "Dial"
	d.App.MajorVersion = "1"
	d.App.MinorVersion = "0"
	d.App.Loaded = true
	errChan := make(chan error)

	// // event.OnceEvent("event::core::loaded", func(event models.Event) {
	// // })

	// // eventsEM.Listen(func(eventType EM.EventType, eventArgs EM.Args) {
	// // 	if eventType == EM.EVENT_DOUBLE_TROUBLE {
	// // 		errChan <- errors.New("DOUBLE TROUBLE, OH SHIT!")
	// // 	} else if eventType == EM.EVENT_STATUS_OK {
	// // 		errChan <- nil
	// // 	}
	// // })

	// event.Push(models.NewEvent("get::core::status"))

	return <-errChan
}

func (d *Dial) Execute(args Args) error {
	// err := EM.Dispatch(&EM.Args{
	// 	Command: "event::app::dial",
	// 	Args:    AppArgs,
	// })

	return nil
}

func (d *Dial) Unload() {
	d.Loaded = false
}
