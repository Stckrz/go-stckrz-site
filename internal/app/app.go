package application

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/stckrz/go-stckrz-site/internal/db"
	"github.com/stckrz/go-stckrz-site/internal/handlers"
	"github.com/stckrz/go-stckrz-site/internal/routes"
	"gorm.io/gorm"
)

type App struct {
	router http.Handler
	db *gorm.DB
}

func New() (*App, error){
	database, err := db.ConnectDB();
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to database: %w", err)
	}

	posts, err := handlers.LoadPosts()
	if err != nil {
		log.Fatal("Failed to load posts", err)
	}
	categories, err := handlers.UniqueCategories(posts)
	if err != nil {
		log.Fatal("Failed to load posts", err)
	}

	app := &App{
		router: routes.LoadRoutes(database, posts, categories),
		db: database,
	}
	return app, nil

}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr: ":3000",
		Handler: a.router,
	}
	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("Failed to start server: %w", err)
	}
	return nil

}
