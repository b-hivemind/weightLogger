package api

import (
	"time"

	"bhavdeep.me/weight_logger/pkg/db"
	"github.com/gin-gonic/gin"
)

func handleGetEntries(c *gin.Context) {
	claims, err := getClaimsFromToken(c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	validator := entriesQuery{}
	if err := c.ShouldBindUri(&validator); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	entries, err := db.WeightByTimeFrame(claims.UUID, validator.Days)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(200, entries)
	}
}

func handleNewEntry(c *gin.Context) {
	claims, err := getClaimsFromToken(c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
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
		Date:   time.Now().Unix(),
		UID:    claims.UUID,
		Weight: validator.Weight,
	}
	err = db.WriteWeight(data)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(200, data)
	}
}
