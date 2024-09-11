package controllers

import (
	"github.com/fajryalvin12/fgh21-go-event-organizer/dtos"
	"github.com/fajryalvin12/fgh21-go-event-organizer/lib"
	"github.com/fajryalvin12/fgh21-go-event-organizer/models"
	"github.com/fajryalvin12/fgh21-go-event-organizer/repository"
	"github.com/gin-gonic/gin"
)

func CreateTransaction(ctx *gin.Context) {
	form := dtos.FormTransaction{}
	ctx.Bind(&form)
	user := ctx.GetInt("userId")

	trx := repository.CreateNewTransactions(models.Transaction{
		EventId: form.EventId,
		PaymentMethodId: form.PaymentMethodId,
		UserId: user,
	})

	for i := range form.SectionId {
		repository.CreateDetailTransaction(models.DetailsTrx{
			SectionId: form.SectionId[i],
			TicketQty: form.TicketQty[i],
			TransactionId: trx.Id,
		})
	}

	data := repository.ListOfTransactions(trx.Id)

	lib.HandlerOk(ctx, "Success to create new transaction", nil, data)
}
func ListOfTransactionsByUserId (ctx *gin.Context) {
	user := ctx.GetInt("userId")

	trx := repository.FindTransactionByUserId(user)
	lib.HandlerOk(ctx, "List user transactions", nil, trx)
}