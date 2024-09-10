package models

type Events struct {
	Id          int    `json:"id" db:"id"`
	Image       string `json:"image" db:"image"`
	Title       string `json:"title" db:"title"`
	Date        string `json:"date" db:"date"`
	Description string `json:"description" db:"description"`
	LocationId  *int   `json:"location_id" db:"location_id"`
	CreatedBy   *int   `json:"created_by" db:"created_by"`
}
type Section struct {
	Id           int    `json:"id"`
	EventId      int    `json:"eventId"`
	SectionName  string `json:"name"`
	Quantity     int    `json:"quantity"`
	SectionPrice int    `json:"price"`
}