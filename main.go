package main

import (
	"database/sql"
	"log"

	"example.com/ris/models"
	"example.com/ris/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	//Open a connection to a PostgreSQL database.
	db, err := sql.Open("postgres", "postgres://postgres:3508@localhost/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	models.CreateTablesinDB(db)

	r := gin.Default()
	routes.InitialiseRoutes(r, db)

	//Start running the server and listen at port 9000.
	if err := r.Run(":9000"); err != nil {
		log.Fatal(err)
	}
}
