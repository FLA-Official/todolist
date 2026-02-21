package userHandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"todolist/model"
	"todolist/utils"
)

// UpdateUser handles PUT /users/{id} and updates an existing user.
func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {

	userID := r.PathValue("id")

	id, err := strconv.Atoi(userID)

	if err != nil {
		http.Error(w, "Please give me a valid user id", http.StatusBadRequest)
		return
	}

	var newUser model.User
	// creating decoder object
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&newUser)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please provide valid json", http.StatusBadRequest)
		return
	}

	newUser.ID = id
	err = h.userrepo.UpdateUser(&newUser)
	if err != nil {
		fmt.Println("UpdateUser error:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	utils.SendData(w, newUser, http.StatusOK)

}
