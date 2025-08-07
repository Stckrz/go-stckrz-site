package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/stckrz/go-stckrz-site/internal/handlers"
)

// handler := &PageHandler{DB: connectd}
func LoadRoutes(posts []handlers.Post, categories []string) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Route("/", func(r chi.Router) {
		LoadPageRoutes(r, posts, categories)
	})
	return router
}

func LoadPageRoutes(router chi.Router, posts []handlers.Post, categories []string) {
	pageHandler := &handlers.PageHandler{Posts: posts, Categories: categories}
	router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("public"))))
	router.Get("/", pageHandler.Index)
	router.Get("/about", pageHandler.About)
	router.Get("/faq", pageHandler.FAQ)
	router.Get("/postpreview", pageHandler.PostPreview)
	router.Get("/postcategorylist", pageHandler.CategoryList)
	router.Get("/posts/{slug}", pageHandler.ViewPost)
}
