package dtos

type Events struct {
	Image       string `form:"image"`
	Title       string `form:"title"`
	Date        string `form:"date"`
	Description string `form:"description"`
	LocationId  *int   `form:"location_id"`
	CreatedBy   *int   `form:"created_by"`
}