package models

import "time"

//User structure
type User struct {
	UserId    int       `json:"userid"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	FavGenres string    `json:"favgenres"`
	CreatedAt time.Time `json:"createdat"`
}
