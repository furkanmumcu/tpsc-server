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
	r.Use(security())

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "it works!!!")
	})

	if os.Getenv("DATABASE_URL") != "" {
		dbase, err := db.OpenDB()
		if err != nil {
			log.Fatalf("Error opening database: %q", err)
		}

		r.GET("/createDB", db.CreateDB(dbase))
		r.GET("/createPassanger/:id/:name/:vehicle/:isOk", db.CreatePassanger(dbase))
		r.GET("/getPassanger/:id/", db.GetPassanger(dbase))
		r.GET("/getCount", db.GetCount(dbase))
		r.GET("/getAllPassangers", db.GetAllPassangers(dbase))
		r.GET("/getVehicle/:vehicle", db.GetVehicle(dbase))
		r.GET("/updateStatus/:id/:status", db.UpdateStatus(dbase))
		r.GET("/updateName/:id/:name", db.UpdateName(dbase))
		r.GET("/updateVehicle/:id/:vehicle", db.UpdateVehicle(dbase))
		r.GET("/deletePassanger/:id/", db.DeletePassanger(dbase))
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "9191"
	}

	r.Run(":" + port)

}

func security() gin.HandlerFunc {
	return func(c *gin.Context) {
		requiredToken := os.Getenv("API_TOKEN")
		if requiredToken == "" {
			requiredToken = "someToken"
		}

		token := c.Request.Header.Get("api_token")
		log.Println("reqToken: " + requiredToken)
		log.Println("recToken: " + token)

		if token == "" || token != requiredToken {
			respondWithError(c, 401, "git burdan :(")
			return
		}

		c.Next()

	}
}

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}
