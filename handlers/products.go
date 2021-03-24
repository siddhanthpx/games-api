package handlers

import (
	"log"
	"net/http"
	"rest_api/data"
)

type Games struct {
	l *log.Logger
}

func NewGames(l *log.Logger) *Games {
	return &Games{l}
}

func (g *Games) GetGames(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetGames()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Couldn't get list of games", http.StatusNotFound)
	}

}
