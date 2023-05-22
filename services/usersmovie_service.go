package services

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type UserMovieService interface {
	AddMovieForUser(ctx *gin.Context, db *sql.DB)
	GetUserMovies(ctx *gin.Context, db *sql.DB)
	DeleteUsersMovie(ctx *gin.Context, db *sql.DB)
}
