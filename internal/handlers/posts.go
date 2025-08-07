package handlers

import (
	"bytes"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/adrg/frontmatter"
	"github.com/go-chi/chi/v5"
	"github.com/gomarkdown/markdown"
)

type metadata struct {
	Title string `yaml:"title"`
	Date  string `yaml:"date"`
	Slug  string `yaml:"slug"`
	Tag   string `yaml:"tag"`
}
type Post struct {
	Title   string
	Date    time.Time
	Slug    string
	Tag     string
	Content template.HTML
}

func (h *PageHandler) PostPreview(w http.ResponseWriter, r *http.Request) {
	h.renderTemplate(w, "partials/postpreview", map[string]any{
		"Title":     "My Blog Posts",
		"Posts": h.Posts,
	})
}

func (h *PageHandler) CategoryList(w http.ResponseWriter, r *http.Request) {
	h.renderTemplate(w, "partials/postcategorylist", map[string]any{
		"Title":      "My Blog Posts",
		"Categories": h.Categories,
	})
}

func (h *PageHandler) ViewPost(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	for _, post := range h.Posts {
		if post.Slug == slug {
			h.renderTemplate(w, "post", map[string]any{
				"Title": post.Title,
				"Post":  post,
			})
			return
		}
	}
	http.NotFound(w, r)
}

func UniqueCategories(posts []Post) ([]string, error) {
	categories := make([]string, 0, len(posts))
	seen := make(map[string]struct{}, len(posts))
	for _, post := range posts {
		if _, exists := seen[post.Tag]; exists {
			continue
		}
		seen[post.Tag] = struct{}{}
		categories = append(categories, post.Tag)
	}

	return categories, nil
}

func LoadPosts() ([]Post, error) {
	dir := "internal/posts"
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var posts []Post
	for _, file := range entries {
		if !strings.HasSuffix(file.Name(), "md") {
			continue
		}

		data, err := os.ReadFile(filepath.Join(dir, file.Name()))
		if err != nil {
			return nil, err
		}

		var meta metadata
		body, err := frontmatter.Parse(bytes.NewReader(data), &meta)
		if err != nil {
			return nil, err
		}

		html := markdown.ToHTML(body, nil, nil)
		rawDate := meta.Date
		parsedDate, err := time.Parse("2006-01-02", rawDate)
		if err != nil {
			return nil, err
		}

		posts = append(posts, Post{
			Title:   meta.Title,
			Date:    parsedDate,
			Slug:    meta.Slug,
			Tag:     meta.Tag,
			Content: template.HTML(html),
		})
	}
	return posts, nil
}
