package main

import (
	"log"

	"github.com/ulicod3/tweeto/bd"
	"github.com/ulicod3/tweeto/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("disconnected to the bd")
		return
	}
	handlers.Handlers()
}
