package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://10.0.0.184:8080")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "bearer, content-type")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func getRouter() *gin.Engine {
	r := gin.Default()
	r.Use(corsMiddleware())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/register", handleRegister)
	r.POST("/login", handleLogin)
	r.Use(jwtMiddleWare())
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
