package userHandler

import (
	"encoding/json"
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

	ctx := r.Context()
	logger := utils.LoggerFromContext(ctx)

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&reqlogin)
	if err != nil {
		logger.Error("invalid request body")
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	usr, err := h.userService.Login(ctx, reqlogin.Email, reqlogin.Password)
	if err != nil {
		logger.Error("login failed", "email", reqlogin.Email)
		http.Error(w, "Invalid Credentials", http.StatusBadRequest)
		return
	}

	cnf := config.GetConfig()

	accessToken, err := utils.CreateJWT(cnf.JWTSecretKey, utils.Payload{
		ID:       usr.ID,
		Username: usr.Username,
		Fullname: usr.Fullname,
		Email:    usr.Email,
	})
	if err != nil {
		logger.Error("failed to create jwt", "user_id", usr.ID)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	logger.Info("user login success", "user_id", usr.ID)

	utils.SendData(w, accessToken, http.StatusCreated)
}
