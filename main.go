package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main(){
	fmt.Println("Hello");
	
	godotenv.Load("config.env") ;

	port := os.Getenv(("PORT"));
	if port == ""{
		log.Fatal("Port was not initialized")
	}

	router := chi.NewRouter();

		router.Use(cors.Handler(cors.Options{
			// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
			AllowedOrigins:   []string{"https://*", "http://*"},
			// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300, // Maximum value not ignored by any of major browsers
		  }))
	srv := &http.Server{
		Handler: router,
		Addr : ":" + port,
	}

	v1Router := chi.NewRouter()

	v1Router.Get("/healthz", handlerReady)
	v1Router.Get("/err", handlerErr)
	router.Mount("/v1", v1Router)

	log.Printf("Running on  Port %v", port );
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

