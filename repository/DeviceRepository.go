package repository

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Device struct {
	IDDev int64  `db:"id_dev"`
	Name  string `db:"name"`
}

type DeviceRepository struct {
	DB *sql.DB
}

func NewDeviceRepository(db *sql.DB) *DeviceRepository {
	return &DeviceRepository{DB: db}
}

func (r *DeviceRepository) GetAllBy() ([]Device, error) {
	query := `
        SELECT d.id_dev, d.name, hlpg.* FROM device d
        LEFT JOIN hl_parking_gate hlpg ON hlpg.id_dev = d.id_dev
        WHERE d.id_reader IS NOT NULL
        AND hlpg.id_dev IS NULL
    `

	rows, err := r.DB.Query(query)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	var devices []Device

	for rows.Next() {
		var device Device
		if err := rows.Scan(&device.IDDev, &device.Name); err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		devices = append(devices, device)
	}

	return devices, nil
}
