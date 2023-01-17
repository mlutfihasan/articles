package main

import (
	"log"
	"net/http"
	"os"

	"app"
	"helper"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	db := app.ConnectDatabase()

	// Validator
	validate := validator.New()
	helper.RegisterValidation(validate)

	router := app.NewRouter(db, validate)
	server := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	log.Printf("Server is running on port %s", port)

	err = server.ListenAndServe()
	helper.PanicIfError(err)
}
