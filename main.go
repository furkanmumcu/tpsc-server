package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"restGin/db"
	"restGin/model"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func dbFuncCreateDB(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//createDb if not exist
		if _, err := db.Exec("CREATE TABLE IF NOT EXISTS passanger (id varchar(255),name varchar(255),vehicle varchar(255), isOk varchar(255), primary key(id))"); err != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("Error creating database table: %q", err))
			return
		}
	}
}

func dbFuncCreatePassanger(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		name := c.Param("name")
		vehicle := c.Param("vehicle")
		isOk := "false"

		fmt.Println(id + name + vehicle + isOk)

		sqlStatement := "INSERT INTO passanger VALUES ($1, $2, $3, $4)"
		if _, err := db.Exec(sqlStatement, id, name, vehicle, isOk); err != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("Error creating database table: %q", err))
			return
		}
		c.String(http.StatusOK, "success")
	}
}

func dbFuncGetPassanger(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		sqlStatement := "SELECT * FROM passanger WHERE id = $1"

		passanger := model.Passanger{}
		row := db.QueryRow(sqlStatement, id)
		err := row.Scan(&passanger.Id, &passanger.Name, &passanger.Vehicle, &passanger.IsOk)

		if err != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("Error %q", err))
			return
		}

		jResult, err := json.Marshal(passanger)
		if err != nil {
			return
		}

		c.String(http.StatusOK, string(jResult))
	}
}

func dbFuncGetCount(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		sqlStatement := "SELECT COUNT(*) FROM passanger"
		row := db.QueryRow(sqlStatement)
		var count int
		err := row.Scan(&count)
		if err != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("Error %q", err))
			return
		}
		c.String(http.StatusOK, strconv.Itoa(count))
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
		db, err := db.OpenDB()
		if err != nil {
			log.Fatalf("Error opening database: %q", err)
		}

		r.GET("/createDB", dbFuncCreateDB(db))
		r.GET("/createPassanger/:id/:name/:vehicle", dbFuncCreatePassanger(db))
		r.GET("/getPassanger/:id/", dbFuncGetPassanger(db))
		r.GET("/getCount", dbFuncGetCount(db))
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "9191"
	}

	r.Run(":" + port)
}
