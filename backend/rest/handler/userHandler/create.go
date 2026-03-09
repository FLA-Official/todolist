package userHandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todolist/model"
	"todolist/utils"
)

// CreateUserHandler handles POST /users and adds a new user to the database.
func (h *Handler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {

	var newUser model.User
	// creating decoder object
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&newUser)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please provide valid json", http.StatusBadRequest)
		return
	}
	err = h.userService.Register(&newUser) // create user
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// creating encoder object
	utils.SendData(w, newUser, http.StatusCreated)

}
