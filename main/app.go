package main

import (
	"log"
	"net/http"

	"github.com/gmurdoko/middleware/handler"
	"github.com/gmurdoko/middleware/middleware"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.Use(middleware.ActivityLogMiddleware)

	r.HandleFunc("/", handler.NewHomeHandler().Handler)
	r.HandleFunc("/product", handler.NewProductHandler().Handler)

	log.Print("Server is listening")
	if err := http.ListenAndServe("localhost:3000", r); err != nil {
		log.Panic(err)
	}
	// r.HandleFunc("/products", ProductsHandler)
	// r.HandleFunc("/articles", ArticlesHandler)
	// http.Handle("/", r)
}
