package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Port not found in environment")
		return
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port not found in environment")
	}

	fmt.Println("Port: ", port)

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("Server staring on port %v", port)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
