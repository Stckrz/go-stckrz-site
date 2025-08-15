package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
	"sort"

)

type PageHandler struct {
	Posts      []Post
	Categories []string
}

func (h *PageHandler) renderTemplate(w http.ResponseWriter, tmpl string, data any) {
	t, err := template.ParseFiles(
		"internal/templates/layout.html",
		"internal/templates/navbar.html",
		"internal/templates/partials/postpreview.html",
		"internal/templates/partials/postcategorylist.html",
		"internal/templates/partials/guestbookform.html",
		"internal/templates/partials/fidgetslider.html",
		filepath.Join("internal/templates", tmpl+".html"),
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.Execute(w, data)
}

func (h *PageHandler) Index(w http.ResponseWriter, r *http.Request) {
	posts := h.Posts

	//Logic to date the posts newest first
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Date.After(posts[j].Date)
	})
	pagedPosts := posts[0:4]

	h.renderTemplate(w, "index", map[string]any{
		"Title":      "Home",
		"Posts":      pagedPosts,
		"Categories": h.Categories,
	})
}

func (h *PageHandler) About(w http.ResponseWriter, r *http.Request) {
	h.renderTemplate(w, "about", map[string]string{
		"Title": "About",
	})
}

func (h *PageHandler) FAQ(w http.ResponseWriter, r *http.Request) {
	h.renderTemplate(w, "faq", map[string]string{"Title": "Faq"})
}
func (h *PageHandler) Resources(w http.ResponseWriter, r *http.Request) {
	h.renderTemplate(w, "resources", map[string]string{"Title": "Resources"})
}
func (h *PageHandler) FidgetSlider(w http.ResponseWriter, r *http.Request) {
	h.renderTemplate(w, "fidgetslider", map[string]string{"Title": "Fidget Slider"})
}
