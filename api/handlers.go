package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/notaryanramani/find-my-bnb/api/store"
	"github.com/notaryanramani/find-my-bnb/api/utils"
	"github.com/notaryanramani/find-my-bnb/api/vectordb"
)

// AuthMiddleware is a middleware that authenticates the user by checking the jwt token
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Authenticating for %s", r.URL.Path)
		token, err := r.Cookie("jwt-token")
		if err != nil {
			http.Error(w, "Unauthorized. No Cookie named \"jwt-token\"", http.StatusUnauthorized)
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

// createUserHandler creates a new user in the database
func (s *Server) createUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received %s request at %s", r.Method, r.URL.Path)

	var userJSON store.UserJSON

	err := json.NewDecoder(r.Body).Decode(&userJSON)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ok := utils.ValidatePassword(userJSON.Password)
	if !ok {
		http.Error(w, "Password must be at least 8 characters long, contains special and uppercase character", http.StatusBadRequest)
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
		if err.Error() == "pq: duplicate key value violates unique constraint \"users_username_key\"" {
			http.Error(w, "Username already exists", http.StatusBadRequest)
			return
		}
		if err.Error() == "pq: duplicate key value violates unique constraint \"users_email_key\"" {
			http.Error(w, "Email already exists", http.StatusBadRequest)
			return
		}
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

// userLoginHandler logs in the user by setting the jwt & username cookie
func (s *Server) userLoginHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received %s request at %s", r.Method, r.URL.Path)

	var userLogin store.UserLogin
	err := json.NewDecoder(r.Body).Decode(&userLogin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := s.store.User.GetByUsername(r.Context(), userLogin.Username)
	if err != nil && err.Error() == "sql: no rows in result set" {
		http.Error(w, "No account found with username: "+userLogin.Username, http.StatusInternalServerError)
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
		MaxAge:   86400,
		HttpOnly: false,
		SameSite: http.SameSiteLaxMode,
		Secure:   false,
	})

	http.SetCookie(w, &http.Cookie{
		Name:  "username",
		Value: user.Username,
		Path:  "/",
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (s *Server) autoLoginHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received %s request at %s", r.Method, r.URL.Path)

	username, err := r.Cookie("username")
	if err != nil {
		http.Error(w, "No username found", http.StatusBadRequest)
		return
	}

	user, err := s.store.User.GetByUsername(r.Context(), username.Value)
	if err != nil && err.Error() == "sql: no rows in result set" {
		http.Error(w, "No account found with username: "+username.Value, http.StatusInternalServerError)
		return
	}

	token, err := utils.GenerateToken(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:   "jwt-token",
		Value:  token,
		Path:   "/",
		MaxAge: 86400,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Auto logged in successfully. Generated new token."))
}

// userLogoutHandler logs out the user by deleting the jwt & username cookie
func (s *Server) userLogoutHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received %s request at %s", r.Method, r.URL.Path)

	http.SetCookie(w, &http.Cookie{
		Name:   "jwt-token",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})

	http.SetCookie(w, &http.Cookie{
		Name:  "username",
		Value: "",
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Logged out successfully"))
}

// getRandomRoomsHandler handles the get random rooms request
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

// getRoomByIdHandler handles the get room by id request
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

// vectorSearchHandler handles the vector search request
func (s *Server) vectorSearchHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received %s request at %s", r.Method, r.URL.Path)

	var vsr vectordb.VectorSearchRequest
	err := json.NewDecoder(r.Body).Decode(&vsr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var nodes []*vectordb.Node
	var query_id string
	var info string

	if _, ok := s.vectordb.ResultCache[vsr.QueryID]; !ok || vsr.QueryID == "" {
		nodes, query_id = s.vectordb.SimilaritySearch(vsr)
		info = "query_id not provided, expired or invalid. New query_id generated."
	} else {
		nodes = s.vectordb.GetNodesFromCache(vsr)
		query_id = vsr.QueryID
		info = "query_id found in cache. Returning cached results."
	}

	ids := make([]int64, len(nodes))
	for i, node := range nodes {
		ids[i] = node.ID
	}

	rooms, err := s.store.Room.GetByMultipleIDs(r.Context(), ids)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	roomsPayload := &store.RoomsPayload{
		Rooms:   rooms,
		QueryID: query_id,
		Info:    info,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(roomsPayload)
}

/* Helper Handlers */

// HelloWord is a simple handler that returns "Hello, World!"
func HelloWord(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received %s request at %s", r.Method, r.URL.Path)
	w.Write([]byte("Hello, World!"))
}

// Check is a simple handler that returns "healthy"
func Check(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received %s request at %s", r.Method, r.URL.Path)
	w.Write([]byte("healthy"))
}

// protectedHandler is a simple handler that returns "protected". Used for testing AuthMiddleware
func (s *Server) protectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("protected"))
}
