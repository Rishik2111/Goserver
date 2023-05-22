package models

import "time"

//Movie stucture
type Movie struct {
	MovieId   int       `json:"movieid"`
	Title     string    `json:"title"`
	Genre     string    `json:"genre"`
	Storyline string    `json:"storyline"`
	CreatedAt time.Time `json:"createdat"`
}
