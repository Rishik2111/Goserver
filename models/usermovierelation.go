package models

//Relationship b/w user and movie along with the rating given by user to that particular movie.
type UserMovieRelation struct {
	UserId  int `json:"userid"`
	MovieId int `json:"movieid"`
	Rating  int `json:"rating"`
}
