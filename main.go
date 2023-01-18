package main

import (
	"fmt"

	"net/http"

	"github.com/mlutfihasan/articles/config"
	"github.com/mlutfihasan/articles/resource"
	"github.com/rs/cors"
)

func main() {
	db := config.Connect()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"POST", "GET"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		Debug:            false,
	})

	mux := new(http.ServeMux)

	mux.Handle("/articles", resource.Articles(db))

	handler := c.Handler(mux)
	server := new(http.Server)
	server.Handler = handler
	server.Addr = ":8080"

	fmt.Println("Starting server at", server.Addr)
	server.ListenAndServe()

}
