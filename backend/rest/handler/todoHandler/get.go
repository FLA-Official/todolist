package todoHandler

import (
	"net/http"
	"todolist/utils"
)

// GetProducts handles GET /products and returns a list of products.
func (h *Handler) GetTasks(w http.ResponseWriter, r *http.Request) {
	// creating encoder object
	utils.SendData(w, h.taskrepo.List(), http.StatusOK)
}
