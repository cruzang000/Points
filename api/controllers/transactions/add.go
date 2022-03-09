package transactions

import (
	"Points/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
	"time"
)

// initialize transactions
var allTransactions []models.Transaction

// AddTransaction acts as endpoint, adds transaction to transactions
func AddTransaction(context *gin.Context) {
	// initialize new transaction
	var newTransaction models.Transaction

	// bind context params to transaction
	err := context.BindJSON(&newTransaction)

	// handle error, return
	if err != nil {
		context.JSON(http.StatusBadRequest, "Error, must Payer, Points and Timestamp must be defined in transaction.")
		return
	}

	// add transaction to transactions
	allTransactions = append(allTransactions, newTransaction)

	// return 201 response
	context.JSON(http.StatusCreated, newTransaction)
}

// AddTransactions takes map of transactions and adds to transactions list
func AddTransactions(transactions []models.PayerSpendPoints) {
	for index := range transactions {
		allTransactions = append(allTransactions, models.Transaction{Payer: transactions[index].Payer, Points: transactions[index].Points, Timestamp: time.Now()})
	}
}

// getTransactions returns sorted transactions
func getTransactions() []models.Transaction {
	sort.Slice(allTransactions, func(index, indexTwo int) bool {
		return allTransactions[index].Timestamp.Before(allTransactions[indexTwo].Timestamp)
	})

	return allTransactions
}
