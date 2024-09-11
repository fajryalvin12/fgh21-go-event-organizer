package dtos

type Events struct {
	Image       string `form:"image"`
	Title       string `form:"title"`
	Date        string `form:"date"`
	Description string `form:"description"`
	LocationId  *int   `form:"location_id"`
	CreatedBy   *int   `json:"created_by" form:"created_by"`
}

type FormSection struct {
	EventId      int    `form:"eventId"`
	SectionName  string `form:"name"`
	Quantity     int    `form:"quantity"`
	SectionPrice int    `form:"price"`
}