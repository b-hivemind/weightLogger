package api

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func newRouter() *mux.Router {
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/", success).Methods("GET")
	myRouter.HandleFunc("/dashboard", stats)
	myRouter.HandleFunc("/new", createNewEntry).Methods("POST", "OPTIONS")
	myRouter.HandleFunc("/all", all)
	return myRouter
}

func HandleRequests() {
	myRouter := newRouter()
	origins := handlers.AllowedOrigins([]string{"*"})
	log.Fatal(http.ListenAndServe(":10000", handlers.CORS(handlers.AllowedHeaders([]string{"content-type"}), origins, handlers.AllowCredentials())(myRouter)))
}
