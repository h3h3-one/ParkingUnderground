package model

type HlInside struct {
	IDCard      string `gorm:"primaryKey;column:id_card" json:"id_card"`
	EnterTime   string `json:"entertime"`
	CounterID   int64  `json:"counterid"`
	ParkingName string `json:"parkingname"`
}

func NewHlInside(idCard, enterTime string, counterID int64, parkingName string) *HlInside {
	return &HlInside{
		IDCard:      idCard,
		EnterTime:   enterTime,
		CounterID:   counterID,
		ParkingName: parkingName,
	}
}
