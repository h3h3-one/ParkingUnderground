package model

type HlParkingGate struct {
	ID        uint   `gorm:"primaryKey;column:id" json:"id"`
	IDDev     int    `gorm:"default:0" json:"id_dev"`
	IDParking int    `json:"id_parking"`
	IDDB      int    `json:"id_db"`
	IsEnter   string `gorm:"default:'1'" json:"is_enter"`
	Name      string `json:"name"`
}

func NewHlParkingGate(id uint, idDev, idParking, idDB int, isEnter, name string) *HlParkingGate {
	return &HlParkingGate{
		ID:        id,
		IDDev:     idDev,
		IDParking: idParking,
		IDDB:      idDB,
		IsEnter:   isEnter,
		Name:      name,
	}
}
