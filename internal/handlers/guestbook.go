package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"time"

	"github.com/stckrz/go-stckrz-site/internal/pagination"
	"gorm.io/gorm"
)

type GuestbookEntry struct {
	gorm.Model
	ID      int
	Date    time.Time
	Name    string
	Email   string
	Website string
	Number  int
	Message string
}

func (h *PageHandler) GuestbookList(w http.ResponseWriter, r *http.Request) {
	currentPage := r.URL.Query().Get("page")
	response, err := http.Get("http://127.0.0.1:8080/guestbooks")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var guestbooks []GuestbookEntry

	if err := json.Unmarshal(responseData, &guestbooks); err != nil {
		fmt.Println(err.Error())
		return
	}
	sort.Slice(guestbooks, func(i, j int) bool {
		return guestbooks[i].CreatedAt.After(guestbooks[j].CreatedAt)
	})

	pageSize := 5
	pagination := pagination.Paginate(guestbooks, currentPage, pageSize)
	h.renderTemplate(w, "guestbooklist", map[string]any{
		"Title":      "Guestbook List",
		"Guestbooks": pagination.Items,
		"CurrentPage": pagination.CurrentPage,
		"TotalPages":  pagination.TotalPages,
		"PageNumbers": pagination.PageNumbers,
	})
}

func (h *PageHandler) GuestbookForm(w http.ResponseWriter, r *http.Request) {
	h.renderTemplate(w, "partials/guestbookform", map[string]any{
		"Title": "Guestbook Form",
	})
}
