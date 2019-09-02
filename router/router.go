package router

import (
	"log"
	"net/http"
	"os"
	"tpsc/db"

	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "it works!!!")
	})

	if os.Getenv("DATABASE_URL") != "" {
		dbase, err := db.OpenDB()
		if err != nil {
			log.Fatalf("Error opening database: %q", err)
		}

		r.GET("/createDB", db.CreateDB(dbase))
		r.GET("/createPassanger/:id/:name/:vehicle", db.CreatePassanger(dbase))
		r.GET("/getPassanger/:id/", db.GetPassanger(dbase))
		r.GET("/getCount", db.GetCount(dbase))
		r.GET("/getAllPassangers", db.GetAllPassangers(dbase))
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "9191"
	}

	r.Run(":" + port)

}
