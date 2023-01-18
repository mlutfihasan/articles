package main

import (
	"fmt"

	"github.com/mlutfihasan/articles/Auth"
	"github.com/mlutfihasan/articles/Config"
	"github.com/mlutfihasan/articles/Resource/General"

	"net/http"

	"github.com/rs/cors"
)

func main() {
	dbPuc := Config.Connect()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"POST", "GET"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		Debug:            false,
	})

	mux := new(Auth.CustomMux)

	mux.Handle("/articles", General.Article(dbPuc))

	handler := c.Handler(mux)
	server := new(http.Server)
	server.Handler = handler
	server.Addr = ":8080"

	fmt.Println("Starting server at", server.Addr)
	server.ListenAndServe()

}
