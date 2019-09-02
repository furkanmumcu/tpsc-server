package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"tpsc/model"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func dummy() {
	fmt.Println("")
}

func CreateDB(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//createDb if not exist
		if _, err := db.Exec("CREATE TABLE IF NOT EXISTS passanger (id varchar(255),name varchar(255),vehicle varchar(255), isOk varchar(255), primary key(id))"); err != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("Error creating database table: %q", err))
			return
		}
	}
}

func CreatePassanger(db *sql.DB) gin.HandlerFunc {
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

func GetPassanger(db *sql.DB) gin.HandlerFunc {
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

func GetCount(db *sql.DB) gin.HandlerFunc {
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

func GetAllPassangers(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		sqlStatement := "SELECT * FROM passanger"
		rows, err := db.Query(sqlStatement)
		if err != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("Error %q", err))
			return
		}

		defer rows.Close()

		var passangers []model.Passanger
		for rows.Next() {
			passanger := model.Passanger{}
			err := rows.Scan(&passanger.Id, &passanger.Name, &passanger.Vehicle, &passanger.IsOk)
			if err != nil {
				c.String(http.StatusInternalServerError,
					fmt.Sprintf("Error %q", err))
				return
			}
			passangers = append(passangers, passanger)
		}

		jResult, err := json.Marshal(passangers)
		if err != nil {
			return
		}

		c.String(http.StatusOK, string(jResult))
	}
}
