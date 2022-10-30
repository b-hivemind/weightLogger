package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func checkAllowedOrigins(c *gin.Context) string {
	ALLOWED_ORIGINS := []string{
		"http://10.0.0.184:8080",
		"http://10.0.0.184:3000",
		"http://localhost:8080",
		"http://localhost:3000",
		"http://10.0.0.228:8080",
	}
	for _, origin := range ALLOWED_ORIGINS {
		if origin == c.Request.Header.Get("Origin") {
			return c.Request.Header.Get("Origin")
		}
	}
	return ""
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", checkAllowedOrigins(c))
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
