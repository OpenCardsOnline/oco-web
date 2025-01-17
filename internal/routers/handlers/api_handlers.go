package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/opencardsonline/oco-web/config"
	"github.com/opencardsonline/oco-web/internal/database"
	"github.com/opencardsonline/oco-web/internal/models"
	"github.com/opencardsonline/oco-web/internal/services"
)

type APIHandlers struct {
	appConfig   *config.AppConfig
	db          *database.AppDBConn
	authService *services.AuthService
}

func (_h *APIHandlers) InitializeAPIHandlers(config *config.AppConfig, db *database.AppDBConn) {
	_h.appConfig = config
	_h.db = db
	_h.authService = &services.AuthService{}
	_h.authService.New(config, db)
}

func (_h *APIHandlers) GetHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func (_h *APIHandlers) AuthRegisterNewUser(w http.ResponseWriter, r *http.Request) {
	var data *models.NewUserRequest

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdUser := _h.authService.CreateNewUser(*data)
	if createdUser == nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("An error occurred when attempting to create the user."))
		return
	}
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(createdUser); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (_h *APIHandlers) AuthVerifyNewUser(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if token != "" {

	} else {
		http.Error(w, "token is required", http.StatusBadRequest)
		return
	}

	isUserVerified := _h.authService.VerifyNewUser(token)
	if isUserVerified {
		w.Write([]byte("<p>User was verified successfully. You may now login!</p>"))
		return
	} else {
		http.Error(w, "user could not be verified", http.StatusInternalServerError)
		return
	}
}
