package main

import (
	"log"
	"net/http"
	"runtime"

	"github.com/quillercs/core/controllers"
	"github.com/quillercs/core/event"
	"github.com/quillercs/core/middleware"
	"github.com/quillercs/core/models"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	eventMachine := event.NewEventMachine()
	eventMachine.Start()

	indexHandler := http.HandlerFunc(controllers.Index)

	http.Handle("/", middleware.CheckEngine(indexHandler))

	eventMachine.OnceEvent("core::status::command", func(event models.Event) {
		ev := models.CreateEvent("core::status::response")
		ev.UUID = event.UUID
		ev.Data = []byte("OK")
		eventMachine.Push(ev)
	})

	for i := 0; i < 100; i++ {
		go func() {
			eventMachine.Command("core::status::command", "core::status::response", func(event models.Event) {
				log.Println("Status", event.Data)
			})
		}()
	}

	log.Println("Running...")
	http.ListenAndServe("127.0.0.1:3000", nil)
}
