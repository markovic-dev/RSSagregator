package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/markovic-dev/rssagg/internal/database"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type paramaters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := paramaters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, "Error parsing JSON")
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		respondWithError(w, 400, "Couldn't create user")
		return
	}

	respondWithJSON(w, 200, databaseUserToUser(user))
}
