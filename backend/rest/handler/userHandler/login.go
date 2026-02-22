package userHandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todolist/config"
	"todolist/utils"
)

// ReqLogin represents the expected JSON body for login requests.
type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Login authenticates a user and returns a JWT access token on success.
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var reqlogin ReqLogin
	// creating decoder object
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&reqlogin)
	if err != nil {
		fmt.Println(err)
		// http.Error(w, "Please provide valid json", 400)
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}
	usr, err := h.userrepo.Find(reqlogin.Email, reqlogin.Password)

	if err != nil {
		http.Error(w, "Invalid Credentials", http.StatusBadRequest)
		return
	}

	cnf := config.GetConfig()
	// JWT secret key is alternatively called access token
	accessToken, err := utils.CreateJWT(cnf.JWTSecretKey, utils.Payload{
		ID:       usr.ID,
		Username: usr.Username,
		Fullname: usr.Fullname,
		Gmail:    usr.Gmail,
	})
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}

	// creating encoder object
	utils.SendData(w, accessToken, http.StatusCreated)
}
