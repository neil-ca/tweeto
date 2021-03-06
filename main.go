package main

import (
	"log"

	"github.com/Neil-uli/tweeto/bd"
	"github.com/Neil-uli/tweeto/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("disconnected to the bd")
		return
	}
	handlers.Handlers()
}
