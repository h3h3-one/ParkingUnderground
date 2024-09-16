package repository

import (
	"database/sql"
	"log"
)

type HlParking struct {
	ID      int64  `db:"id"`
	IDDB    int64  `db:"id_db"`
	Name    string `db:"name"`
	Enabled int64  `db:"enabled"`
}

type ParkingRepository struct {
	db *sql.DB
}

func NewParkingRepository(db *sql.DB) *ParkingRepository {
	return &ParkingRepository{db: db}
}

func (repo *ParkingRepository) FindAll() ([]HlParking, error) {
	rows, err := repo.db.Query("SELECT * FROM HL_PARKING")
	if err != nil {
		log.Println("Error retrieving all parking:", err)
		return nil, err
	}
	defer rows.Close()

	var parkings []HlParking
	for rows.Next() {
		var parking HlParking
		if err := rows.Scan(&parking.ID, &parking.IDDB, &parking.Name, &parking.Enabled); err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		parkings = append(parkings, parking)
	}
	return parkings, nil
}

func (repo *ParkingRepository) FindByID(id int64) (*HlParking, error) {
	var parking HlParking
	err := repo.db.QueryRow("SELECT * FROM HL_PARKING WHERE ID = ?", id).Scan(&parking.ID, &parking.IDDB, &parking.Name, &parking.Enabled)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Парковка не найдена
		}
		log.Println("Error finding parking by ID:", err)
		return nil, err
	}
	return &parking, nil
}

func (repo *ParkingRepository) UpdateParking(parking HlParking) error {
	query := "UPDATE HL_PARKING SET ID_DB = 1, NAME = ?, ENABLED = ? WHERE ID = ?"
	_, err := repo.db.Exec(query, parking.Name, parking.Enabled, parking.ID)
	if err != nil {
		log.Println("Error updating parking:", err)
		return err
	}
	return nil
}

func (repo *ParkingRepository) DeleteByID(id int64) error {
	_, err := repo.db.Exec("DELETE FROM HL_PARKING WHERE ID = ?", id)
	if err != nil {
		log.Println("Error deleting parking by ID:", err)
		return err
	}
	return nil
}

func (repo *ParkingRepository) InsertByName(name string) error {
	query := "INSERT INTO HL_PARKING (ID_DB, NAME, ENABLED) VALUES (1, ?, 1)"
	_, err := repo.db.Exec(query, name)
	if err != nil {
		log.Println("Error inserting parking by name:", err)
		return err
	}
	return nil
}
