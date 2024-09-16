package event

type NewEvent struct {
	IsEvent bool `json:"isEvent"`
}

func NewNewEvent() *NewEvent {
	return &NewEvent{
		IsEvent: false,
	}
}
