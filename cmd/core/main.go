package main

import (
	"log"
	"net/http"
	"runtime"

	"github.com/quillercs/core/event"

	"github.com/quillercs/core/controllers"
	"github.com/quillercs/core/middleware"
)

type Core struct {
	eventMachine *event.Machine
}

func NewCore() *Core {
	return &Core{}
}

func (c *Core) Start() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	c.eventMachine = event.NewMachine()

	indexHandler := http.HandlerFunc(controllers.Index)

	http.Handle("/", middleware.CheckEngine(indexHandler))

	log.Println("Running...")
	http.ListenAndServe("127.0.0.1:3000", nil)
}

func main() {
	c := NewCore()
	c.Start()
}
