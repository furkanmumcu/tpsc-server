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
		isOk := c.Param("isOk")

		fmt.Println(id + name + vehicle + isOk)

		sqlStatement := "INSERT INTO passanger VALUES ($1, $2, $3, $4)"
		_, err := db.Exec(sqlStatement, id, name, vehicle, isOk)
		handleError(c, err)
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
		handleError(c, err)

		jResult, err := json.Marshal(passanger)
		handleError(c, err)

		c.String(http.StatusOK, string(jResult))
	}
}

func GetCount(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		sqlStatement := "SELECT COUNT(*) FROM passanger"
		row := db.QueryRow(sqlStatement)
		var count int
		err := row.Scan(&count)
		handleError(c, err)
		c.String(http.StatusOK, strconv.Itoa(count))
	}
}

func GetAllPassangers(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		sqlStatement := "SELECT * FROM passanger"
		rows, err := db.Query(sqlStatement)
		handleError(c, err)

		defer rows.Close()

		var passangers []model.Passanger
		for rows.Next() {
			passanger := model.Passanger{}
			err := rows.Scan(&passanger.Id, &passanger.Name, &passanger.Vehicle, &passanger.IsOk)
			handleError(c, err)
			passangers = append(passangers, passanger)
		}

		jResult, err := json.Marshal(passangers)
		handleError(c, err)

		c.String(http.StatusOK, string(jResult))
	}
}

func GetVehicle(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		vehicle := c.Param("vehicle")
		sqlStatement := "SELECT * FROM passanger where vehicle = $1"
		rows, err := db.Query(sqlStatement, vehicle)
		handleError(c, err)

		defer rows.Close()

		var passangers []model.Passanger
		for rows.Next() {
			passanger := model.Passanger{}
			err := rows.Scan(&passanger.Id, &passanger.Name, &passanger.Vehicle, &passanger.IsOk)
			handleError(c, err)
			passangers = append(passangers, passanger)
		}

		jResult, err := json.Marshal(passangers)
		handleError(c, err)

		c.String(http.StatusOK, string(jResult))
	}
}

func UpdateStatus(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		status := c.Param("status")
		sqlStatement := "UPDATE passanger SET isOk = $1 where id = $2"
		_, err := db.Exec(sqlStatement, status, id)
		handleError(c, err)

		c.String(http.StatusOK, "updated status")
	}
}

func UpdateName(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		name := c.Param("name")
		sqlStatement := "UPDATE passanger SET name = $1 where id = $2"
		_, err := db.Exec(sqlStatement, name, id)
		handleError(c, err)

		c.String(http.StatusOK, "updated name")
	}
}

func UpdateVehicle(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		status := c.Param("status")
		sqlStatement := "UPDATE passanger SET vehicle = $1 where id = $2"
		_, err := db.Exec(sqlStatement, status, id)
		handleError(c, err)

		c.String(http.StatusOK, "updated vehicle")
	}
}

func handleError(c *gin.Context, err error) {
	if err != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprintf("Error %q", err))
		return
	}
}
