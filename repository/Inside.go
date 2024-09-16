package repository

import (
	"database/sql"
	"log"
)

type HlInside struct {
	CounterID   string `db:"counterid"`
	Place       string `db:"place"`
	ParkingName string `db:"parkingname"`
}

type InsideRepository struct {
	DB *sql.DB
}

func NewInsideRepository(db *sql.DB) *InsideRepository {
	return &InsideRepository{DB: db}
}

func (r *InsideRepository) SelectAllInside(parkingID int64) ([]HlInside, error) {
	query := `
        SELECT hli.*, hlc.name as place, hlp.name as parkingname 
        FROM hl_inside hli
        JOIN hl_counters hlc ON hlc.id = hli.counterid
        JOIN hl_parking hlp ON hlp.id = hli.counterid 
        WHERE hlp.id = $1
    `

	rows, err := r.DB.Query(query, parkingID)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	var insides []HlInside

	for rows.Next() {
		var inside HlInside
		if err := rows.Scan(&inside.CounterID, &inside.Place, &inside.ParkingName); err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		insides = append(insides, inside)
	}

	return insides, nil
}
