package services

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type MovieServices interface {
	CreateMovie(ctx *gin.Context, db *sql.DB)
	GetMovie(ctx *gin.Context, db *sql.DB)
	GetAllMovies(ctx *gin.Context, db *sql.DB)
	// UpdateUser(ctx *gin.Context, db *sql.DB)
	DeleteMovie(ctx *gin.Context, db *sql.DB)
}
