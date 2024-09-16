package repository

import (
	"database/sql"
)

type HlParkingGate struct {
	IDParking int64  `db:"id_parking"`
	IDDev     int64  `db:"id_dev"`
	IDDB      int64  `db:"id_db"`
	IsEnter   string `db:"is_enter"`
	Name      string `db:"name"`
}

type ParkingGateRepository struct {
	db *sql.DB
}

func NewParkingGateRepository(db *sql.DB) *ParkingGateRepository {
	return &ParkingGateRepository{db: db}
}

func (repo *ParkingGateRepository) UpdateGate(idParking, idDev int64, isEnter string) error {
	query := "UPDATE HL_PARKING_GATE SET IS_ENTER = ? WHERE ID_PARKING = ? AND ID_DEV = ? AND ID_DB = 1"
	_, err := repo.db.Exec(query, isEnter, idParking, idDev)
	return err
}

func (repo *ParkingGateRepository) InsertGate(idParking int64, name, isEnter string) error {
	query := "INSERT INTO HL_PARKING_GATE (ID_PARKING, ID_DEV, ID_DB, IS_ENTER, NAME) VALUES (?, ?, 1, ?, ?)"
	_, err := repo.db.Exec(query, idParking, 0, isEnter, name)
	return err
}

func (repo *ParkingGateRepository) GetAllGate(id int64) ([]HlParkingGate, error) {
	rows, err := repo.db.Query("SELECT hlp.id_parking, hlp.id_dev, hlp.id_db, hlp.is_enter, d.name, hlp.ID FROM HL_PARKING_GATE hlp JOIN device d ON d.id_dev=hlp.id_dev WHERE hlp.id_parking = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var gates []HlParkingGate
	for rows.Next() {
		var gate HlParkingGate
		if err := rows.Scan(&gate.IDParking, &gate.IDDev, &gate.IDDB, &gate.IsEnter, &gate.Name); err != nil {
			return nil, err
		}
		gates = append(gates, gate)
	}
	return gates, nil
}

func (repo *ParkingGateRepository) GetDeviceByID(id int64) (*HlParkingGate, error) {
	var gate HlParkingGate
	err := repo.db.QueryRow("SELECT hlp.id_parking, hlp.id_dev, hlp.id_db, hlp.is_enter, d.name FROM HL_PARKING_GATE hlp JOIN device d ON d.id_dev=hlp.id_dev WHERE hlp.id_dev = ?", id).Scan(&gate.IDParking, &gate.IDDev, &gate.IDDB, &gate.IsEnter, &gate.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &gate, nil
}

func (repo *ParkingGateRepository) DeleteByID(id int64) error {
	_, err := repo.db.Exec("DELETE FROM HL_PARKING_GATE WHERE ID_DEV = ?", id)
	return err
}
