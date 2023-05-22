package services

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"example.com/ris/models"
	"github.com/gin-gonic/gin"
)

type Info struct {
	MovieId   int       `json:"movieid"`
	Title     string    `json:"title"`
	Genre     string    `json:"genre"`
	Storyline string    `json:"storyline"`
	CreatedAt time.Time `json:"createdat"`
	Rating    int       `json:"rating"`
}

func AddMovieForUser(ctx *gin.Context, db *sql.DB) {
	var relation models.UserMovieRelation
	if err := ctx.ShouldBindJSON(&relation); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `INSERT INTO usermovierelation (userid, movieid, rating) VALUES ($1, $2, $3)`
	db.QueryRow(query, relation.UserId, relation.MovieId, relation.Rating)

	ctx.JSON(http.StatusCreated, relation)
}

func GetUserMovies(ctx *gin.Context, db *sql.DB) {
	userId, err := strconv.Atoi(ctx.Param("User_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}

	query := `SELECT movies.movieid, movies.title, movies.genre, movies.storyline, movies.createdat, usermovierelation.rating FROM usermovierelation INNER JOIN movies ON usermovierelation.movieid = movies.movieid WHERE usermovierelation.userid = $1`
	rows, err := db.Query(query, userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	Movies := []Info{}
	for rows.Next() {
		var movie Info
		if err := rows.Scan(&movie.MovieId, &movie.Title, &movie.Genre, &movie.Storyline, &movie.CreatedAt, &movie.Rating); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		Movies = append(Movies, movie)
	}

	ctx.JSON(http.StatusOK, Movies)
}

func DeleteUsersMovie(ctx *gin.Context, db *sql.DB) {
	userID := ctx.Query("userid")
	movieID := ctx.Query("movieid")

	result, err := db.Exec("DELETE FROM usermovierelation WHERE userid = $1 AND movieid = $2", userID, movieID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	CntRowsAffected, err := result.RowsAffected()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if CntRowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Movie not found in user list"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Movie deleted successfully from the user list"})
}
