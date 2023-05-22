package services

import (
	"database/sql"
	"net/http"
	"strconv"

	"example.com/ris/models"
	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context, db *sql.DB) {

	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `INSERT INTO users (name, email, favgenres) VALUES ($1, $2, $3) RETURNING userid, createdat`
	err := db.QueryRow(query, user.Name, user.Email, user.FavGenres).Scan(&user.UserId, &user.CreatedAt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

func GetUser(ctx *gin.Context, db *sql.DB) {
	userID, err := strconv.Atoi(ctx.Param("User_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}

	var user models.User
	query := `SELECT userid, name, email, favgenres, createdat FROM users WHERE userid = $1`
	err = db.QueryRow(query, userID).Scan(&user.UserId, &user.Name, &user.Email, &user.FavGenres, &user.CreatedAt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func GetAllUsers(ctx *gin.Context, db *sql.DB) {

	query := `SELECT userid, name, email, favgenres, createdat FROM users`
	rows, err := db.Query(query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	Users := []models.User{}
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.UserId, &user.Name, &user.Email, &user.FavGenres, &user.CreatedAt); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		Users = append(Users, user)
	}

	ctx.JSON(http.StatusOK, Users)
}

func DeleteUser(ctx *gin.Context, db *sql.DB) {
	userID, err := strconv.Atoi(ctx.Param("User_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}

	result, err := db.Exec("DELETE FROM users WHERE userid = $1", userID)
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

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
