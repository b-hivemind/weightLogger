package api

import (
	"time"

	"bhavdeep.me/weight_logger/pkg/db"
	"github.com/gin-gonic/gin"
)

func convertToKG(entries []db.Entry) []db.Entry {
	var convertedEntries []db.Entry
	for _, entry := range entries {
		var convertedEntry db.Entry
		convertedEntry.Weight = 0.45 * entry.Weight
		convertedEntry.Date = entry.Date
		convertedEntries = append(convertedEntries, convertedEntry)
	}
	return convertedEntries
}

func handleGetEntries(c *gin.Context) {
	units := c.DefaultQuery("units", "LB")
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
		if units == "KG" {
			entries = convertToKG(entries)
		}
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
