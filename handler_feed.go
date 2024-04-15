package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"


	"github.com/elBanna00/rss-agg/internal/database"
	"github.com/google/uuid"
)

func (apiCnfg *apiConfig) handlerCreateFeed(w http.ResponseWriter , r *http.Request, user database.User){
	type parameters struct {
		Name string `json:"name"`
		Url string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w ,400 ,fmt.Sprintf("Error parsing JSON %v", err))
		return
	}
	feed ,err := apiCnfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID : uuid.New(),
		CreatedAt : time.Now().UTC(),
		UpdatedAt : time.Now().UTC(),
		Name : params.Name,
		Url : params.Url,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w ,400 ,fmt.Sprintf("Couldn't Create feed %v", err))
		return
	}
	respondWithJSON(w , 201 ,databaseFeedToFeed(feed))
}
func (apiCnfg *apiConfig) handlerGetFeeds(w http.ResponseWriter , r *http.Request ){
	feeds ,err := apiCnfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w ,400 ,fmt.Sprintf("Couldn't find feeds %v", err))
		return
	}
	respondWithJSON(w , 201 ,databaseFeedsToFeeds(feeds))
}