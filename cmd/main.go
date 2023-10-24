package main

import (
	"log"
	"refactoring/internal/app"
)

func main() {

	if err := app.Run(); err != nil {
		log.Fatal("failed to run server: %v", err)
	}

}
