package model

type HlParking struct {
	ID      uint   `gorm:"primaryKey;column:id" json:"id"`
	IDDB    int64  `json:"id_db"`
	Name    string `json:"name"`
	Enabled int64  `json:"enabled"`
}

func NewHlParking(id uint, idDB int64, name string, enabled int64) *HlParking {
	return &HlParking{
		ID:      id,
		IDDB:    idDB,
		Name:    name,
		Enabled: enabled,
	}
}
