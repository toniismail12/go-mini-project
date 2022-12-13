package main

import (
	"Tugas-16/controllers"
	"Tugas-16/database"
	"database/sql"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("Faild load file env")
	} else {
		fmt.Println("Success read file env")
	}

	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"),
	)

	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()
	if err != nil {
		fmt.Println("Connection Failed")
		panic(err)
	} else {
		fmt.Println("Connection Success")
	}

	database.DbMigrate(DB)

	defer DB.Close()

	router := gin.Default()
	router.GET("/persons", controllers.GetAllPerson)
	router.POST("/persons", controllers.InsertPerson)
	router.PUT("/persons/:id", controllers.UpdatePerson)
	router.DELETE("/persons/:id", controllers.DeletePerson)

	router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
