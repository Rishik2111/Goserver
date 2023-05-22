package services

import (
	"database/sql"
	"net/http"
	"strconv"

	"example.com/ris/models"
	"github.com/gin-gonic/gin"
)

func CreateMovie(ctx *gin.Context, db *sql.DB) {

	var movie models.Movie
	if err := ctx.ShouldBindJSON(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `INSERT INTO movies (title, genre, storyline) VALUES ($1, $2, $3) RETURNING movieid, createdat`
	err := db.QueryRow(query, movie.Title, movie.Genre, movie.Storyline).Scan(&movie.MovieId, &movie.CreatedAt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, movie)
}

func GetMovie(ctx *gin.Context, db *sql.DB) {
	movieID, err := strconv.Atoi(ctx.Param("Movie_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Movie ID"})
		return
	}

	var movie models.Movie
	query := `SELECT movieid, title, genre, storyline, createdat FROM movies WHERE movieid = $1`
	err = db.QueryRow(query, movieID).Scan(&movie.MovieId, &movie.Title, &movie.Genre, &movie.Storyline, &movie.CreatedAt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, movie)
}

func GetAllMovies(ctx *gin.Context, db *sql.DB) {

	query := `SELECT movieid, title, genre, storyline, createdat FROM movies`
	rows, err := db.Query(query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	Movies := []models.Movie{}
	for rows.Next() {
		var movie models.Movie
		if err := rows.Scan(&movie.MovieId, &movie.Title, &movie.Genre, &movie.Storyline, &movie.CreatedAt); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		Movies = append(Movies, movie)
	}

	ctx.JSON(http.StatusOK, Movies)
}

func DeleteMovie(ctx *gin.Context, db *sql.DB) {
	movieID, err := strconv.Atoi(ctx.Param("Movie_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Movie ID"})
		return
	}

	result, err := db.Exec("DELETE FROM movies WHERE movieid = $1", movieID)
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
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Movie deleted successfully"})
}
