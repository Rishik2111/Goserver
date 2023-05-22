package routes

import (
	"database/sql"

	"example.com/ris/services"
	"github.com/gin-gonic/gin"
)

// Routes associated with corresponding HTTP method and appropriate service function from the services package.
func InitialiseRoutes(r *gin.Engine, db *sql.DB) {

	// User Routes
	r.POST("/users/createuser", func(ctx *gin.Context) {
		services.CreateUser(ctx, db)
	})
	r.GET("/users/getall", func(ctx *gin.Context) {
		services.GetAllUsers(ctx, db)
	})
	r.GET("/users/getuser/:User_id", func(ctx *gin.Context) {
		services.GetUser(ctx, db)
	})
	r.DELETE("/users/deleteuser/:User_id", func(ctx *gin.Context) {
		services.DeleteUser(ctx, db)
	})

	// Movie Routes
	r.POST("/movies/createmovie", func(ctx *gin.Context) {
		services.CreateMovie(ctx, db)
	})
	r.GET("/movies/getall", func(ctx *gin.Context) {
		services.GetAllMovies(ctx, db)
	})
	r.GET("/movies/getmovie/:Movie_id", func(ctx *gin.Context) {
		services.GetMovie(ctx, db)
	})
	r.DELETE("/movies/deletemovie/:Movie_id", func(ctx *gin.Context) {
		services.DeleteMovie(ctx, db)
	})

	// User-Movie Relation Routes
	r.POST("/usermovie/add", func(ctx *gin.Context) {
		services.AddMovieForUser(ctx, db)
	})
	r.GET("/usermovie/getusermovies/:User_id", func(ctx *gin.Context) {
		services.GetUserMovies(ctx, db)
	})
	r.DELETE("/usermovie/deleteusermovie", func(ctx *gin.Context) {
		services.DeleteUsersMovie(ctx, db)
	})
}
