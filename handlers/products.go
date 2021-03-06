package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"rest_api/cache"
	"rest_api/data"
	"strconv"

	"github.com/gorilla/mux"
)

var postCache cache.PostCache

type Games struct {
	l *log.Logger
}

func NewGames(l *log.Logger) *Games {
	return &Games{l}
}

func (g *Games) GetGames(rw http.ResponseWriter, r *http.Request) {
	g.l.Println("Handle GET Request")
	lp := data.GetGames()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Couldn't get list of games", http.StatusNotFound)
	}

}

func (g *Games) GetGame(rw http.ResponseWriter, r *http.Request) {
	g.l.Println("Handle GET Request")
	vars := mux.Vars(r)
	var post *data.Game = postCache.Get(vars["id"])
	id, err := strconv.Atoi(vars["id"])

	if post == nil {
		game, err := data.GetGame(id)

		if err != nil {
			http.Error(rw, "Cannot parse ID", http.StatusBadRequest)
			return
		}

		postCache.Set(vars["id"], game)
		enc := json.NewEncoder(rw)
		enc.Encode(game)
	} else {
		enc := json.NewEncoder(rw)
		enc.Encode(post)

	}

	if err != nil {
		http.Error(rw, "Cannot find game with ID", http.StatusNotFound)
	}

}

func (g *Games) AddGame(rw http.ResponseWriter, r *http.Request) {
	g.l.Println("Handle POST Request")
	game := data.NewGame()
	d := json.NewDecoder(r.Body)
	err := d.Decode(game)
	if err != nil {
		http.Error(rw, "Cannot parse JSON", http.StatusBadRequest)
	}

	data.AddGame(game)
	e := json.NewEncoder(rw)
	err = e.Encode(game)
	if err != nil {
		http.Error(rw, "Cannot present JSON", http.StatusBadRequest)
	}

}
