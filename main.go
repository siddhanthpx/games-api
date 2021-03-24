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

	getRouter := r.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", gl.GetGames)
	r.Use(loggingMiddleware)

	// Create Server
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(rw, r)
	})
}
