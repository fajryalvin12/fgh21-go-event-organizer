package models

type Wishlist struct {
	Id 		int 	`json:"id"`
	UserId 	int		`json:"userId" form:"userId" db:"user_id"`
	EventId int 	`json:"eventId" form:"eventId" db:"event_id"`
}
type JoinWishlistEvent struct {
	Id 			int 	`json:"id"`
	Title 		string	`json:"title"`
	Date 		string	`json:"date"`
	Location 	*int	`json:"location"`
	Description string 	`json:"description"`
}