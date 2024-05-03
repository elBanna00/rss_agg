package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/elBanna00/rss-agg/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCnfg *apiConfig) handlerCreateFollowedFeed(w http.ResponseWriter , r *http.Request, user database.User){
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`

	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w ,400 ,fmt.Sprintf("Error parsing JSON %v", err))
		return
	}
	feedFollow ,err := apiCnfg.DB.CreateFollowedFeed(r.Context(), database.CreateFollowedFeedParams{
		ID : uuid.New(),
		CreatedAt : time.Now().UTC(),
		UpdatedAt : time.Now().UTC(),
		FeedID: params.FeedID,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w ,400 ,fmt.Sprintf("Couldn't Create feed follow %v", err))
		return
	}
	respondWithJSON(w , 201 ,databaseFeedFollowToFeedFollow(feedFollow))
}
func (apiCnfg *apiConfig) handlerGetFollowedFeed(w http.ResponseWriter , r *http.Request, user database.User){
	feedFollows ,err := apiCnfg.DB.GetFollowedFeed(r.Context(), user.ID)
	if err != nil {
		respondWithError(w ,400 ,fmt.Sprintf("Couldn't Get Followed Feeds %v", err))
		return
	}
	respondWithJSON(w , 201 ,databaseFeedFollowsToFeedFollows(feedFollows))
}

func (apiCnfg *apiConfig) handlerDeleteFollowedFeed(w http.ResponseWriter , r *http.Request, user database.User){
	followedFeedstr := chi.URLParam(r , "followedFeedID")
	followedFeedID , err := uuid.Parse(followedFeedstr)
	if err != nil {
		
		respondWithError(w ,400 ,fmt.Sprintf("Couldn't Parse Followed Feed ID %v", err))
		return
	}
	err = apiCnfg.DB.DeleteFollowedFeed(r.Context(), database.DeleteFollowedFeedParams{
		ID : followedFeedID,
		UserID: user.ID,
	})

	if err != nil {
		
		respondWithError(w ,400 ,fmt.Sprintf("Couldn't Delete Followed Feeds %v", err))
		return
	}
	respondWithJSON(w,200, struct{}{})
}