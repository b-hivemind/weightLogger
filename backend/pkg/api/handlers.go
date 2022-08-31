package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"bhavdeep.me/weight_logger/pkg/db"
	"github.com/gin-gonic/gin"
)

func handleEntries(c *gin.Context) {
	validator := entriesQuery{}
	if err := c.ShouldBindUri(&validator); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	entries, err := db.WeightByTimeFrame(validator.days)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(200, entries)
	}
}

func stats(w http.ResponseWriter, r *http.Request) {
	var response [][]db.Entry
	// Last 2 days, to calculate delta
	entries, err := db.WeightByTimeFrame(2)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	response = append(response, entries)
	entries, err = db.WeightByTimeFrame(7)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	response = append(response, entries)
	entries, err = db.WeightByTimeFrame(30)
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
	data := db.Entry{
		Date:   time.Now().Format("2006-01-02"),
		Weight: string(response.Weight),
	}
	if !response.Force {
		lastEntry, err := db.WeightByTimeFrame(1)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
		if len(lastEntry) > 0 && lastEntry[0].Date == data.Date {
			w.WriteHeader(http.StatusMultipleChoices)
			w.Write([]byte("300"))
			return
		}
	}
	err = db.WriteWeight(data, response.Force)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	json.NewEncoder(w).Encode(data)
}
