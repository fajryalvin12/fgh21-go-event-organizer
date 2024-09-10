package models

type Transaction struct {
	Id int 					`json:"id"`
	EventId int 			`json:"eventId" db:"event_id"`
	PaymentMethodId int 	`json:"paymentMethodId" db:"payment_method_id"`
	UserId int 				`json:"userId" db:"user_id"`
}
type DetailTransaction struct {
	TransactionId int 		`json:"transactionId" db:"transaction_id"`
	FullName string 		`json:"fullName" db:"full_name"`
	EventTitle string 		`json:"eventTitle" db:"event_title"`
	LocationId *int 			`json:"locationId" db:"location_id"`
	Date string 			`json:"date"`
	PaymentMethod string 	`json:"paymentMethod" db:"payment_method"`
	TicketSection []string 	`json:"section" db:"ticket_section"`
	Quantity []int 			`json:"ticketQty" db:"quantity"`
}