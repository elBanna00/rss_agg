package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/elBanna00/rss-agg/internal/auth"
	"github.com/elBanna00/rss-agg/internal/database"
	"github.com/google/uuid"
)

func (apiCnfg *apiConfig) handlerCreateUser(w http.ResponseWriter , r *http.Request){
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w ,400 ,fmt.Sprintf("Error parsing JSON %v", err))
		return
	}
	usr ,err := apiCnfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID : uuid.New(),
		CreatedAt : time.Now().UTC(),
		UpdatedAt : time.Now().UTC(),
		Name : params.Name,
	})
	if err != nil {
		respondWithError(w ,400 ,fmt.Sprintf("Couldn't Create user %v", err))
		return
	}
	respondWithJSON(w , 201 ,databaseUsertoUser(usr))
}

func (apiCnfg *apiConfig) handlerGetUserByAPIKey(w http.ResponseWriter , r *http.Request){
	apikey, err := auth.GetApiKey(r.Header);
	if err != nil {
		respondWithError(w, 403, fmt.Sprintf("Auth Error: %v", err))
		return
	}

	user, err := apiCnfg.DB.GetUserByAPIKey(r.Context(), apikey)
	if err != nil {
		respondWithError(w, 400 , fmt.Sprintf("Couldn't get user: %v" , err))
		return
	}
	respondWithJSON(w,200 , databaseUsertoUser(user))

}