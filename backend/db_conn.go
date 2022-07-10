package main

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

var (
	db    *sql.DB
	table = "main"
)

func writeWeight(ent Entry, force bool) error {
	fmt.Println(ent)
	query := fmt.Sprintf("INSERT INTO %s (date, weight) VALUES ('%s', %s)", table, ent.Date, ent.Weight)
	if force {
		query = fmt.Sprintf("UPDATE IGNORE %s SET weight='%s' WHERE date='%s'", table, ent.Weight, ent.Date)
	}
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("writeWeight: %v", err)
	}
	return nil
}

func weightByTimeFrame(days int) ([]Entry, error) {
	var entries []Entry
	query := fmt.Sprintf("SELECT * FROM %s", table)
	if days > 0 {
		query += fmt.Sprintf(" ORDER BY date DESC LIMIT %d", days)
	}
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

func connect() error {
	cfg := mysql.Config{
		User:   "archimedes",
		Passwd: "@r(h1m3d3s",
		Net:    "tcp",
		Addr:   "database:3306",
		DBName: "weight_data",
	}
	//Get a database handle
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return fmt.Errorf("Connect: %s", err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return fmt.Errorf("Connect: %s", pingErr)
	}
	fmt.Println("Connected!")
	return nil
}
