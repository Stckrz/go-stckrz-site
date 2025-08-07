package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/stckrz/go-stckrz-site/internal/handlers"
	"gorm.io/gorm"
)

// handler := &PageHandler{DB: connectd}
func LoadRoutes(database *gorm.DB, posts []handlers.Post, categories []string) *chi.Mux{
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Route("/", func(r chi.Router){
		LoadPageRoutes(r, database, posts, categories)
	})
	return router
}

func LoadPageRoutes(router chi.Router, database *gorm.DB, posts []handlers.Post, categories []string) {
	pageHandler := &handlers.PageHandler{DB: database, Posts: posts, Categories: categories}
	router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("public"))))
	router.Handle("/media/*", http.StripPrefix("/media/", http.FileServer(http.Dir("public/music"))))
	router.Get("/", pageHandler.Index)
	router.Get("/about", pageHandler.About)
	router.Get("/faq", pageHandler.FAQ)
	router.Get("/music", pageHandler.MusicList)
	router.Post("/addsong", pageHandler.AddSong)
	router.Get("/musicplayer/{filename}", pageHandler.MusicPlayer)
	router.Get("/postpreview", pageHandler.PostPreview)
	router.Get("/postcategorylist", pageHandler.CategoryList)
	router.Get("/posts/{slug}", pageHandler.ViewPost)
}
