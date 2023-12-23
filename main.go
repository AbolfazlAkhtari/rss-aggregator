package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello World")

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
}
