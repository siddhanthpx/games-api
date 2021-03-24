package main

import (
	"log"
	"net/http"
	"os"
	"rest_api/handlers"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	l := log.New(os.Stdout, "intro", log.LstdFlags)

	gl := handlers.NewGames(l)

	getAllRouter := r.Methods(http.MethodGet).Subrouter()
	getAllRouter.HandleFunc("/", gl.GetGames)

	getRouter := r.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/{id:[0-9]+}", gl.GetGame)

	postRouter := r.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", gl.AddGame)

	// Create Server
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
