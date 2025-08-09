package handlers

import (
	"fmt"
	"net/http"
	"time"
)

type GuestbookEntry struct {
	Date    time.Time
	Name    string
	Email   string
	Website string
	Number  int
	Message string
}

func (h *PageHandler) GuestbookList(w http.ResponseWriter, r *http.Request) {
	h.renderTemplate(w, "guestbooklist", map[string]any{
		"Title": "Guestbook List",
	})

	// selectedTag := r.URL.Query().Get("tag")
	// currentPage := r.URL.Query().Get("page")
	// pageSize := 5
	//
	// posts := h.Posts
	// if selectedTag != "" {
	// 	filtered := make([]Post, 0, len(posts))
	// 	for _, post := range posts {
	// 		if post.Tag == selectedTag {
	// 			filtered = append(filtered, post)
	// 		}
	// 	}
	// 	posts = filtered
	// }
	//
	// //Logic to date the posts newest first
	// sort.Slice(posts, func(i, j int) bool {
	// 	return posts[i].Date.After(posts[j].Date)
	// })
	//
	// pagination := pagination.Paginate(posts, currentPage, pageSize)
	//
	// h.renderTemplate(w, "postlist", map[string]any{
	// 	"Title":       "Post List",
	// 	"Posts":       pagination.Items,
	// 	"Categories":  h.Categories,
	// 	"SelectedTag": selectedTag,
	// 	"CurrentPage": pagination.CurrentPage,
	// 	"TotalPages":  pagination.TotalPages,
	// 	"PageNumbers": pagination.PageNumbers,
	// })
}

func (h *PageHandler) GuestbookForm(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hii")
	h.renderTemplate(w, "partials/guestbookform", map[string]any{
		"Title": "Guestbook Form",
	})
}
