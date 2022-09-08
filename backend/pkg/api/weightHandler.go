package api

import (
	"fmt"
	"time"

	"bhavdeep.me/weight_logger/pkg/db"
	"github.com/gin-gonic/gin"
)

func handleGetEntries(c *gin.Context) {
	validator := entriesQuery{}
	if err := c.ShouldBindUri(&validator); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	entries, err := db.WeightByTimeFrame(validator.Days)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(200, entries)
	}
}

func handleNewEntry(c *gin.Context) {
	validator := newEntryQuery{}
	if err := c.ShouldBindJSON(&validator); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	if validator.Weight <= 0 {
		c.JSON(400, gin.H{"msg": "Weight cannot be zero or negative"})
		return
	}
	data := db.Entry{
		Date:   time.Now().Format("2006-01-02"),
		Weight: validator.Weight,
	}
	if !validator.Force {
		lastEntry, err := db.WeightByTimeFrame(1)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		if len(lastEntry) > 0 && lastEntry[0].Date == data.Date {
			c.JSON(300, gin.H{"msg": fmt.Sprintf("Weight already exists for %v, use force: true to override", data.Date)})
			return
		}
	}
	err := db.WriteWeight(data, validator.Force)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(200, data)
	}
}
