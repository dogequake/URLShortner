package handler

import (
	"URLShortner/model"
	"URLShortner/service"
	"encoding/json"
	"net/http"
	"strings"
)

type Handler struct {
	Service *service.URLService
}

func NewHandler(s *service.URLService) *Handler {
	return &Handler{Service: s}
}

func (h *Handler) ShortenURL(w http.ResponseWriter, r *http.Request) {
	var req model.URLRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	if err := h.Service.ValidateURL(req.URL); err != nil {
		http.Error(w, "invalid url", http.StatusBadRequest)
		return
	}
	shortID, err := h.Service.SaveURL(req.URL)
	if err != nil {
		http.Error(w, "could not generate short url", http.StatusInternalServerError)
		return
	}
	shortURL := model.ShortURL{ShortURL: "http://147.45.187.208:8080/" + shortID}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(shortURL)
}

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	shortID := strings.TrimPrefix(r.URL.Path, "/")
	if shortID == "" {
		http.NotFound(w, r)
		return
	}
	longURL, ok := h.Service.GetLongURL(shortID)
	if !ok {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, longURL, http.StatusMovedPermanently)
}
