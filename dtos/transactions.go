package dtos

type FormTransaction struct {
	EventId 		int `form:"eventId"`
	PaymentMethodId int `form:"paymentMethodId"`
	TicketQty 		[]int `form:"ticketQty"`
	SectionId 		[]int `form:"sectionId"`	
}