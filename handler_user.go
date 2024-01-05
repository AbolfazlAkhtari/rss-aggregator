package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"gituhb.com/AbolfazlAkhtari/RSS/internal/auth"
	"gituhb.com/AbolfazlAkhtari/RSS/internal/database"
	"net/http"
	"time"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `name:""`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing request json: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Coudn't create a new user: %v", err))
		return
	}

	respondWithJson(w, 201, dbUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	apikey, err := auth.GetApiKey(r.Header)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Auth err: %v", err))
		return
	}

	user, err := apiCfg.DB.GetUserByApiKey(r.Context(), apikey)
	if err != nil {
		respondWithError(w, 404, fmt.Sprintf("Couldn't get User: %v", err))
		return
	}

	respondWithJson(w, 200, dbUserToUser(user))
}
