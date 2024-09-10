package controllers

import (
	"fmt"

	"github.com/fajryalvin12/fgh21-go-event-organizer/lib"
	"github.com/fajryalvin12/fgh21-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

type FormTransaction struct {
	EventId 		int `form:"eventId"`
	PaymentMethodId int `form:"paymentMethodId"`
	TicketQty 		[]int `form:"ticketQty"`
	SectionId 		[]int `form:"sectionId"`	
}
func CreateTransaction(ctx *gin.Context) {
	form := FormTransaction{}
	ctx.Bind(&form)
	user := ctx.GetInt("userId")
	fmt.Println(form)

	trx := models.CreateNewTransactions(models.Transaction{
		EventId: form.EventId,
		PaymentMethodId: form.PaymentMethodId,
		UserId: user,
	})

	for i := range form.SectionId {
		models.CreateDetailTransaction(models.DetailsTrx{
			SectionId: form.SectionId[i],
			TicketQty: form.TicketQty[i],
			TransactionId: trx.Id,
		})
	}

	data := models.ListOfTransactions(trx.Id)

	lib.HandlerOk(ctx, "Success to create new transaction", nil, data)
}
func ListOfTransactionsByUserId (ctx *gin.Context) {
	user := ctx.GetInt("userId")

	trx := models.FindTransactionByUserId(user)
	lib.HandlerOk(ctx, "List user transactions", nil, trx)

}