package db

import (
	"fmt"
)

func WriteWeight(ent Entry, force bool) error {
	query := fmt.Sprintf("INSERT INTO %s (date, weight) VALUES ('%s', %f)", weightLogTable, ent.Date, ent.Weight)
	if force {
		query = fmt.Sprintf("UPDATE IGNORE %s SET weight='%f' WHERE date='%s'", weightLogTable, ent.Weight, ent.Date)
	}
	db, err := connect()
	if err != nil {
		return fmt.Errorf("WriteWeight | Error connecting to database | %v", err)
	}
	defer db.Close()
	_, err = db.Exec(query)
	if err != nil {
		return fmt.Errorf("WriteWeight: %v", err)
	}
	return nil
}

func WeightByTimeFrame(days int) ([]Entry, error) {
	var entries []Entry
	query := fmt.Sprintf("SELECT * FROM %s", weightLogTable)
	if days > 0 {
		query += fmt.Sprintf(" ORDER BY date DESC LIMIT %d", days)
	}
	fmt.Printf("Query: %s\n", query)
	db, err := connect()
	if err != nil {
		return nil, fmt.Errorf("WeightByTimeFrame | Error connecting to database | %v", err)
	}
	defer db.Close()
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("weightByTimeFrame %d: %v", days, err)
	}
	defer rows.Close()
	//Loop through the rows, using Scan to assign column data to struct fields
	for rows.Next() {
		var ent Entry
		if err := rows.Scan(&ent.Date, &ent.Weight); err != nil {
			return nil, fmt.Errorf("weightByTimeFrame %d: %v", days, err)
		}
		entries = append(entries, ent)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("weightByTimeFrame %d: %v", days, err)
	}
	return entries, nil
}
