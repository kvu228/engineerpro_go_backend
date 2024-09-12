package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"simple_server/models"
	"simple_server/storage"
	"simple_server/utils"
)

// RegisterHandler handles user registration
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.MethodNotAllowed(w)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		utils.BadRequest(w, "Username and password are required")
		return
	}

	storage.Mutex.Lock()
	defer storage.Mutex.Unlock()

	if _, exists := storage.Users[username]; exists {
		utils.Conflict(w, "User already exists")
		return
	}

	storage.Users[username] = models.User{
		Username: username,
		Password: password,
	}

	if err := storage.SaveUsers(); err != nil {
		utils.InternalServerError(w, "Failed to save user data")
		return
	}

	utils.OK(w, "User registered successfully")
}

// LoginHandler handles user login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.MethodNotAllowed(w)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	storage.Mutex.Lock()
	user, exists := storage.Users[username]
	storage.Mutex.Unlock()

	if !exists || user.Password != password {
		utils.Unauthorized(w, "Invalid username or password")
		return
	}

	utils.OK(w, "Login successful")
}

// ProfileHandler handles viewing and editing of user profiles
func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		viewProfileHandler(w, r)
	case http.MethodPost:
		editProfileHandler(w, r)
	default:
		utils.MethodNotAllowed(w)
	}
}

func viewProfileHandler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")

	storage.Mutex.Lock()
	user, exists := storage.Users[username]
	storage.Mutex.Unlock()

	if !exists {
		utils.NotFound(w, "User not found")
		return
	}

	json.NewEncoder(w).Encode(user)
}

func editProfileHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	file, _, err := r.FormFile("profile_picture")

	if err != nil && !errors.Is(err, http.ErrMissingFile) {
		utils.BadRequest(w, "Error reading profile picture")
		return
	}

	storage.Mutex.Lock()
	user, exists := storage.Users[username]
	storage.Mutex.Unlock()

	if !exists || user.Password != password {
		utils.Unauthorized(w, "Invalid username or password")
		return
	}

	if file != nil {
		defer file.Close()

		imagePath := filepath.Join("uploads", fmt.Sprintf("%s_profile.png", username))
		imgData, err := io.ReadAll(file)
		if err != nil {
			utils.InternalServerError(w, "Error saving profile picture")
			return
		}
		if err := os.WriteFile(imagePath, imgData, 0644); err != nil {
			utils.InternalServerError(w, "Error saving profile picture")
			return
		}
		user.ProfilePicUrl = imagePath
	}

	storage.Mutex.Lock()
	storage.Users[username] = user
	storage.Mutex.Unlock()

	if err := storage.SaveUsers(); err != nil {
		utils.InternalServerError(w, "Failed to save user data")
		return
	}

	utils.OK(w, "User profile updated successfully")
}
