package models

type Events struct {
	Id          int    `json:"id"`
	Image       string `json:"image"`
	Title       string `json:"title"`
	Date        string `json:"date"`
	Description string `json:"description"`
	LocationId  *int   `json:"location_id"`
	CreatedBy   *int   `json:"created_by" db:"created_by"`
}
type Section struct {
	Id           int    `json:"id"`
	EventId      int    `json:"eventId"`
	SectionName  string `json:"name"`
	Quantity     int    `json:"quantity"`
	SectionPrice int    `json:"price"`
}

type EventLocation struct {
	Id          int    `json:"id"`
	Image       string `json:"image"`
	Title       string `json:"title"`
	Date        string `json:"date"`
	Description string `json:"description"`
	Location    string `json:"location"`
	CreatedBy   *int   `json:"created_by" db:"created_by"`
}