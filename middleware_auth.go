package main

import (
	"fmt"
	"net/http"

	"github.com/elBanna00/rss-agg/internal/auth"
	"github.com/elBanna00/rss-agg/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (cnfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	apikey, err := auth.GetApiKey(r.Header);
	if err != nil {
		respondWithError(w, 403, fmt.Sprintf("Auth Error: %v", err))
		return
	}

	user, err := cnfg.DB.GetUserByAPIKey(r.Context(), apikey)
	if err != nil {
		respondWithError(w, 400 , fmt.Sprintf("Couldn't get user: %v" , err))
		return
	}
	handler(w,r,user)
	}

}