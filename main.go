package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func dbFuncCreate(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//createDb if not exist
		if _, err := db.Exec("CREATE TABLE IF NOT EXISTS passanger (id varchar(255),name varchar(255),vehicle varchar(255), isOk varchar(255))"); err != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("Error creating database table: %q", err))
			return
		}
	}
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		fmt.Println("hello called")
		c.String(http.StatusOK, "hello world")
	})

	r.GET("/user/:name", func(c *gin.Context) {
		fmt.Println("called")
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{"user": name})
	})

	if os.Getenv("DATABASE_URL") != "" {
		db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
		if err != nil {
			log.Fatalf("Error opening database: %q", err)
		}

		r.GET("/createDB", dbFuncCreate(db))
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "9191"
	}

	r.Run(":" + port)
}
