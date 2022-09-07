package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
func newRouter() *mux.Router {
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/", handleIndex).Methods("GET")
	//myRouter.HandleFunc("/days", )
	//myRouter.HandleFunc("/new", createNewEntry).Methods("POST", "OPTIONS")
	return myRouter
}
*/

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func getRouter() *gin.Engine {
	r := gin.Default()
	r.Use(CORSMiddleware())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/entries/:numdays", handleGetEntries)
	r.POST("/entries/new", handleNewEntry)
	return r
}

func HandleRequests() {
	// TOOD probably should use an ENV variable here
	router := getRouter()
	s := &http.Server{
		Addr:    ":10080",
		Handler: router,
	}
	log.Fatal(s.ListenAndServe())
}
