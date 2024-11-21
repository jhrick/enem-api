package handler

import "github.com/go-chi/chi/v5"

type Handler struct {
  Router *chi.Mux
}

func NewHandler() *Handler {
  return &Handler{
    Router: chi.NewMux(),
  }
}
