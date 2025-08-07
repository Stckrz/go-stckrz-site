package handlers

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/stckrz/go-stckrz-site/internal/models"
)
func (h *PageHandler) MusicPlayer(w http.ResponseWriter, r *http.Request) {
	filename := chi.URLParam(r, "filename")
	var song models.Song
	result := h.DB.Where("file_name = ?", filename).First(&song)
	if result.Error != nil {
		http.Error(w, "Cannot find filename..", http.StatusNotFound)
		return
	}
	h.renderTemplate(w, "musicviews/musicplayer", song)
}

func (h *PageHandler) MusicList(w http.ResponseWriter, r *http.Request) {
	var songs []models.Song
	h.DB.Find(&songs)

	h.renderTemplate(w, "musicviews/musiclist", map[string]any{
		"Title": "Music List",
		"Songs": songs,
	})
}

func (h *PageHandler) AddSong(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	song := models.Song{
		FileName: r.Form.Get("filename"),
		Title:    r.Form.Get("title"),
		Artist:   r.Form.Get("artist"),
		Album:    r.Form.Get("album"),
	}
	h.DB.Create(&song)
	http.Redirect(w, r, "/music", http.StatusSeeOther)
}
