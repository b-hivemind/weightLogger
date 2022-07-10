package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// TODO use rs/cors

type Entry struct {
	Date   string `json:"date"`
	Weight string `json:"weight"`
}

type Response_New struct {
	Weight string `json:"weight"`
	Force  bool   `json:"force"`
}

func all(w http.ResponseWriter, r *http.Request) {
	entries, err := weightByTimeFrame(0)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(entries)
}

func success(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "weight_weight_logger_api v0.2.0 :)\n")
}

func stats(w http.ResponseWriter, r *http.Request) {
	var response [][]Entry
	// Last 2 days, to calculate delta
	entries, err := weightByTimeFrame(2)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	response = append(response, entries)
	entries, err = weightByTimeFrame(7)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	response = append(response, entries)
	entries, err = weightByTimeFrame(30)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	response = append(response, entries)
	json.NewEncoder(w).Encode(response)
}

func createNewEntry(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var response Response_New
	if err := json.Unmarshal(reqBody, &response); err != nil {
		fmt.Println(response)
		fmt.Println(string(reqBody))
		fmt.Println(err)
		return
	}
	float_weight, err := strconv.ParseFloat(string(response.Weight), 32)
	if err != nil {
		http.Error(w, "Weight cannot be parsed to float", http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	if float_weight <= 0 {
		http.Error(w, "Invalid weight", http.StatusBadRequest)
		return
	}
	data := Entry{
		Date:   time.Now().Format("2006-01-02"),
		Weight: string(response.Weight),
	}
	if !response.Force {
		lastEntry, err := weightByTimeFrame(1)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
		if lastEntry[0].Date == data.Date {
			w.WriteHeader(http.StatusMultipleChoices)
			w.Write([]byte("300"))
			return
		}
	}
	err = writeWeight(data, response.Force)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	json.NewEncoder(w).Encode(data)
}

func newRouter() *mux.Router {
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/", success).Methods("GET", "OPTIONS")
	myRouter.HandleFunc("/dashboard", stats)
	myRouter.HandleFunc("/new", createNewEntry).Methods("POST", "OPTIONS")
	myRouter.HandleFunc("/all", all)
	return myRouter
}

func handleRequests() {
	myRouter := newRouter()
	origins := handlers.AllowedOrigins([]string{"*"})
	log.Fatal(http.ListenAndServe(":10000", handlers.CORS(handlers.AllowedHeaders([]string{"content-type"}), origins, handlers.AllowCredentials())(myRouter)))
}

func main() {
	connect()
	handleRequests()
	fmt.Println("Now listening on 10000")
}
