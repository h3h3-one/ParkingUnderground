package repository

import (
	"database/sql"
	"log"
)

type HlParam struct {
	TabloIP   string
	TabloPort int64
	BoxIP     string
	BoxPort   int64
	IDGate    int64
	IDCam     int64
	IDDev     int64
}

type ParamRepository struct {
	db *sql.DB
}

func NewParamRepository(db *sql.DB) *ParamRepository {
	return &ParamRepository{db: db}
}

func (repo *ParamRepository) AddParamDevice(param HlParam) error {
	query := `
        INSERT INTO HL_PARAM (TABLO_IP, TABLO_PORT, BOX_IP, BOX_PORT, ID_GATE, ID_CAM, ID_DEV)
        VALUES (?, ?, ?, ?, ?, ?, ?)`

	_, err := repo.db.Exec(query, param.TabloIP, param.TabloPort, param.BoxIP, param.BoxPort, param.IDGate, param.IDCam, param.IDDev)
	if err != nil {
		log.Printf("Error adding param device: %v", err)
		return err
	}
	return nil
}
