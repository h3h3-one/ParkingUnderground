package model

type Device struct {
	ID   uint   `gorm:"primaryKey;column:id_dev" json:"id_dev"`
	Name string `json:"name"`
}

func NewDevice(id uint, name string) *Device {
	return &Device{
		ID:   id,
		Name: name,
	}
}
