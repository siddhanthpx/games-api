package data

import (
	"encoding/json"
	"fmt"
	"io"
)

//   Game Model
type Game struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
}

type Games []*Game

func GetGames() Games {
	return OurGames
}

func GetGame(id int) (*Game, error) {
	g, err := findGame(id)
	return g, err
}

func AddGame(g *Game) {
	g.ID = getNextID()
	OurGames = append(OurGames, g)
}

func (g *Games) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(g)
}

var ErrGameNotFound = fmt.Errorf("Game not found")

func findGame(id int) (*Game, error) {
	for _, g := range OurGames {
		if g.ID == id {
			return g, nil
		}
	}

	return nil, ErrGameNotFound
}

func NewGame() *Game {
	var game Game
	game.ID = getNextID()
	return &game
}

func getNextID() int {
	lp := OurGames[len(OurGames)-1]
	return lp.ID + 1
}

var OurGames = []*Game{
	{
		ID:          1,
		Name:        "Dark Souls",
		Description: "Challenging and thrilling. First in the series.",
		Price:       39.99,
		SKU:         "sca212",
	},

	{
		ID:          2,
		Name:        "Dark Souls 2",
		Description: "A fan favorite. Managing to get a lot of things right.",
		Price:       39.99,
		SKU:         "pwf083",
	},

	{
		ID:          3,
		Name:        "Bloodborne",
		Description: "Set in a vast world where you set out to hunt beasts.",
		Price:       59.99,
		SKU:         "yha197",
	},
}
