package api

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
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

func getRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/entries/:numdays", handleEntries)
	return r
}

func HandleRequests() {
	// TOOD probably should use an ENV variable here
	router := getRouter()
	s := &http.Server{
		Addr:    ":10090",
		Handler: router,
	}
	log.Fatal(s.ListenAndServe())
	// log.Fatal(http.ListenAndServe(":10000", handlers.CORS(handlers.AllowedHeaders([]string{"content-type"}), origins, handlers.AllowCredentials())(myRouter)))
}
