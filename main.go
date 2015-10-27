package main

import (
	"log"
	"net/http"
	"runtime"

	"github.com/quillercs/core/controllers"
	"github.com/quillercs/core/middleware"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	indexHandler := http.HandlerFunc(controllers.Index)

	http.Handle("/", middleware.CheckEngine(indexHandler))

	log.Println("Running...")
	http.ListenAndServe("127.0.0.1:3000", nil)
}
