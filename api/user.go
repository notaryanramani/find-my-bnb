package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"regexp"
)

type UserJSON struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserPayload struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

func (s *Server) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user UserJSON
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Json Decode Error", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id`

	err = s.db.QueryRowContext(
		ctx,
		query,
		user.Username,
		user.Email,
		user.Password,
	).Scan(&user.ID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response := UserPayload{
		ID:       user.ID,
		Username: user.Username,
	}
	json.NewEncoder(w).Encode(response)
}

func validateUserCredentials(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var user UserJSON
		if err := json.Unmarshal(body, &user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if !validateUsername(user.Username) {
			http.Error(w, "Invalid Username", http.StatusBadRequest)
			return
		}

		if !validatePassword(user.Password) {
			http.Error(w, "Invalid Password", http.StatusBadRequest)
			return
		}

		r.Body = io.NopCloser(bytes.NewBuffer(body))
		next.ServeHTTP(w, r)
	}
}

func validateUsername(username string) bool {
	return len(username) >= 5
}

func validatePassword(password string) bool {
	if password == "" {
		return false
	}

	if len(password) < 8 {
		return false
	}

	upperCaseRegex := regexp.MustCompile(`[A-Z]`)
	lowerCaseRegex := regexp.MustCompile(`[a-z]`)
	numberRegex := regexp.MustCompile(`[0-9]`)

	if !upperCaseRegex.MatchString(password) {
		return false
	}

	if !lowerCaseRegex.MatchString(password) {
		return false
	}

	if !numberRegex.MatchString(password) {
		return false
	}
	return true
}
