package repository

import "food/model"

type Guess struct {
	DB model.DataBase
}

func (g *Guess) ListGuess() []model.Guess {
	var guessList []model.Guess
	g.DB.MyDB.Find(&guessList)
	return guessList
}
