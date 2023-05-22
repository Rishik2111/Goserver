package services

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type UserServices interface {
	CreateUser(ctx *gin.Context, db *sql.DB)
	GetUser(ctx *gin.Context, db *sql.DB)
	GetAllUsers(ctx *gin.Context, db *sql.DB)
	// UpdateUser(ctx *gin.Context, db *sql.DB)
	DeleteUser(ctx *gin.Context, db *sql.DB)
}
