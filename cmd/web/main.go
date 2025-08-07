package main

import (
	"context"
	"fmt"
	"log"

	application "github.com/stckrz/go-stckrz-site/internal/app"
	// "github.com/stckrz/go-stckrz-site/internal/handlers"
)


func main() {
	app, err := application.New()
	if err != nil {
		log.Fatal("Failed to initialize application", err)
	}

	if err := app.Start(context.TODO()); err != nil {
		fmt.Println("Failed to start app", err)
	}
}
