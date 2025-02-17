package main

import (
	"encoding/json"
	"net/http"

	"github.com/notaryanramani/find-my-bnb/api/store"
	"github.com/notaryanramani/find-my-bnb/api/utils"
)

func HelloWord(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func Check(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("healthy"))
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Unauthorized. Token Missing", http.StatusUnauthorized)
			return
		}

		tokenString = tokenString[len("Bearer "):]
		err := utils.VerifyToken(tokenString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}

func (s *Server) createUserHandler(w http.ResponseWriter, r *http.Request) {
	var userJSON store.UserJSON

	err := json.NewDecoder(r.Body).Decode(&userJSON)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ok := utils.ValidatePassword(userJSON.Password)
	if !ok {
		http.Error(w, "Password must be at least 8 characters long", http.StatusBadRequest)
		return
	}

	hashedPassword, err := utils.HashPassword(userJSON.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user := &store.User{
		Username: userJSON.Username,
		Email:    userJSON.Email,
		Password: hashedPassword,
	}

	err = s.store.User.Create(r.Context(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := &store.UserPayload{
		ID:       user.ID,
		Username: user.Username,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (s *Server) userLoginHandler(w http.ResponseWriter, r *http.Request) {
	var userLogin store.UserLogin
	err := json.NewDecoder(r.Body).Decode(&userLogin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := s.store.User.GetByUsername(r.Context(), userLogin.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = utils.CompareHash(user.Password, userLogin.Password)
	if err != nil {
		http.Error(w, "Invalid Password", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateToken(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := &store.UserLoginPayload{
		Username: user.Username,
		Token:    token,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
