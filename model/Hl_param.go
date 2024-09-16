package model

type HlParam struct {
	ID        uint   `gorm:"primaryKey;column:id" json:"id"`
	TabloIP   string `json:"tablo_ip"`
	TabloPort string `json:"tablo_port"`
	BoxIP     string `json:"box_ip"`
	BoxPort   string `json:"box_port"`
	IDGate    int64  `json:"id_gate"`
	IDCam     int64  `json:"id_cam"`
	IDDev     int64  `json:"id_dev"`
}

func NewHlParam(id uint, tabloIP, tabloPort, boxIP, boxPort string, idGate, idCam, idDev int64) *HlParam {
	return &HlParam{
		ID:        id,
		TabloIP:   tabloIP,
		TabloPort: tabloPort,
		BoxIP:     boxIP,
		BoxPort:   boxPort,
		IDGate:    idGate,
		IDCam:     idCam,
		IDDev:     idDev,
	}
}
