package transactions

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetBalances acts as api endpoint for getting balances, loops through allTransactions and calcaulates current balances
func GetBalances(context *gin.Context) {
	payerBalances := make(map[string]int)
	transactions := getTransactions()

	// loop transactions and calculate balances
	for _, transaction := range transactions {
		if _, ok := payerBalances[transaction.Payer]; ok {
			payerBalances[transaction.Payer] += transaction.Points
			continue
		}

		payerBalances[transaction.Payer] = transaction.Points
	}

	// return 201 response and balances payload
	context.JSON(http.StatusAccepted, payerBalances)
}
