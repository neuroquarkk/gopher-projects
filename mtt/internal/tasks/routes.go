package tasks

import "net/http"

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /tasks", h.Create)
	mux.HandleFunc("GET /tasks/{id}", h.GetById)
	mux.HandleFunc("DELETE /tasks/{id}", h.Delete)
}
