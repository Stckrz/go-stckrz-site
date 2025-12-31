package application

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	// "github.com/stckrz/go-stckrz-site/internal/config"
	"github.com/stckrz/go-stckrz-site/internal/handlers"
	"github.com/stckrz/go-stckrz-site/internal/routes"
)

type App struct {
	router http.Handler
	// config *config.Config
}

func New() (*App, error) {
	// err := config.Load("config.json")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	posts, err := handlers.LoadPosts()
	if err != nil {
		return nil, fmt.Errorf("failed to load posts: %w", err)
	}
	categories, err := handlers.UniqueCategories(posts)
	if err != nil {
		return nil, fmt.Errorf("failed to load categories: %w", err)
	}

	app := &App{
		router: routes.LoadRoutes(posts, categories),
	}
	return app, nil

}

func (a *App) Start(ctx context.Context) error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := "0.0.0.0:" + port
	log.Printf("Listening on %s\n", addr)

	server := &http.Server{
		Addr:    addr,
		Handler: a.router,
	}

	return server.ListenAndServe()
	// server := &http.Server{
	// 	Addr: ":3000",
	// 	Handler: a.router,
	// }
	// err := server.ListenAndServe()
	// if err != nil {
	// 	return fmt.Errorf("Failed to start server: %w", err)
	// }
	// return nil
	//
}
