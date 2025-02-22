package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
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
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("jwt-token")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenString := token.Value
		err = utils.VerifyToken(tokenString)
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
		Name:     userJSON.Name,
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
		Username: user.Username,
		Message:  "User created successfully",
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

	http.SetCookie(w, &http.Cookie{
		Name:     "jwt-token",
		Value:    token,
		Path:     "/",
		Expires:  time.Now().Add(24 * time.Hour),
		MaxAge:   86400,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (s *Server) getRandomRoomsHandler(w http.ResponseWriter, r *http.Request) {
	var topK store.TopKPayload
	var err error
	err = json.NewDecoder(r.Body).Decode(&topK)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var rooms []*store.RoomPayload
	if topK.Ids == nil {
		rooms, err = s.store.Room.GetTopKRandom(r.Context(), topK.K)
	} else if len(topK.Ids) == 0 {
		w.Header().Set("Warning", "Empty Ids array. Results may contain duplicates.")
		rooms, err = s.store.Room.GetTopKRandom(r.Context(), topK.K)
	} else {
		rooms, err = s.store.Room.NextTopKRandom(r.Context(), topK.K, topK.Ids)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	roomsPayload := &store.RoomsPayload{
		Rooms: rooms,
	}

	json.NewEncoder(w).Encode(roomsPayload)
}

func (s *Server) getRoomByIdHandler(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id64 := int64(id)
	room, err := s.store.Room.GetByID(r.Context(), id64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(room)
}

func (s *Server) protectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Protected."))
}
