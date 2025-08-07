package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"sort"
	"strconv"

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
		filepath.Join("internal/templates", tmpl+".html"),
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.Execute(w, data)
}

func (h *PageHandler) Index(w http.ResponseWriter, r *http.Request) {
	selectedTag := r.URL.Query().Get("tag")
	currentPage := r.URL.Query().Get("page")
	posts := h.Posts
	if currentPage == "" {
		currentPage = "1"
	}

	

	parsedPage, err := strconv.Atoi(currentPage)
	if err != nil {
		fmt.Println("could not parse page")
		return
	}
	if selectedTag != "" {
		filtered := make([]Post, 0, len(posts))
		for _, post := range posts {
			if post.Tag == selectedTag {
				filtered = append(filtered, post)
			}
		}
		posts = filtered
	}

	//Logic to date the posts newest first
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Date.After(posts[j].Date)
	})

	//Pagination information 2
	skip := 2


	totalPages := (len(posts) + skip - 1) / skip
	pageNumbers := make([]int, totalPages)
	for i := range pageNumbers {
		pageNumbers[i] = i + 1
	}

	start := (parsedPage - 1) * skip
	end :=parsedPage * skip
	if start >= len(posts) {
		start = len(posts)
	}
	if end > len(posts) {
		end = len(posts)
	}

	pagedPosts := posts[start:end]

	h.renderTemplate(w, "index", map[string]any{
		"Title":       "Home",
		"Posts":       pagedPosts,
		"Categories":  h.Categories,
		"SelectedTag": selectedTag,
		"CurrentPage": parsedPage,
		"TotalPages": totalPages,
		"PageNumbers": pageNumbers,
	})
}

func (h *PageHandler) About(w http.ResponseWriter, r *http.Request) {
	h.renderTemplate(w, "about", map[string]string{"Title": "About"})
}

func (h *PageHandler) FAQ(w http.ResponseWriter, r *http.Request) {
	h.renderTemplate(w, "faq", map[string]string{"Title": "Faq"})
}
