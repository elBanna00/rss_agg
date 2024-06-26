package main

import (
	// "fmt"
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/elBanna00/rss-agg/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)
type apiConfig struct {
	DB *database.Queries
}



func main(){
//==================================
//Top level code 
//==================================

//Loading env variables
	godotenv.Load("config.env") ;

	port := os.Getenv(("PORT"));
	if port == ""{
		log.Fatal("Port was not initialized")
	}
	dbURL := os.Getenv(("DB_URL"));
	if dbURL == ""{
		log.Fatal("DataBase URL was not initialized")
	}

//===============================================
//Data base init
//==============================================

//opening connection with DB


	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}


	dbQueries := database.New(db)

	apiCnfg := apiConfig{
		DB : dbQueries,
	}

//===============================================
//Creating Routers 
//==============================================

	router := chi.NewRouter();
	v1Router := chi.NewRouter();


		router.Use(cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300, // Maximum value not ignored by any of major browsers
		  }))
//Creating new server
	srv := &http.Server{
		Handler: router,
		Addr : ":" + port,
	}

//EndPoints 

	v1Router.Get("/healthz", handlerReady)
	v1Router.Get("/err", handlerErr)
	v1Router.Post("/users", apiCnfg.handlerCreateUser)
	v1Router.Get("/users", apiCnfg.middlewareAuth(apiCnfg.handlerGetUserByAPIKey))

	v1Router.Post("/feeds", apiCnfg.middlewareAuth(apiCnfg.handlerCreateFeed))

	
	v1Router.Get("/feeds", apiCnfg.handlerGetFeeds)

	v1Router.Post("/followfeeds", apiCnfg.middlewareAuth(apiCnfg.handlerCreateFollowedFeed))

	v1Router.Get("/followfeeds", apiCnfg.middlewareAuth(apiCnfg.handlerGetFollowedFeed))

	v1Router.Delete("/followfeeds/{followedFeedID}", apiCnfg.middlewareAuth(apiCnfg.handlerDeleteFollowedFeed))
	router.Mount("/v1", v1Router)

	log.Printf("Running on  Port %v", port );
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

