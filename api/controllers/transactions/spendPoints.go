package transactions

import (
	"Points/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SpendPoints acts as api endpoint, accepts data matching SpendPointsInput, goes through transactions and calculates
// how much each payer would spend using oldest points first and not leaving any payer points as negative
func SpendPoints(context *gin.Context) {
	// initialize new transaction
	var pointsToSpend models.SpendPointsInput

	// bind context params to transaction
	err := context.BindJSON(&pointsToSpend)
	// handle error, return
	if err != nil || pointsToSpend.Points < 1 {
		context.JSON(http.StatusBadRequest, "Error, must send points greater than 0 ex:{\"points\": 1}.")
		return
	}

	// call function to consolidate transactions, return a payer balance ascending array
	var balancesAscending = consolidateTransactions()

	// call consolidate transactions to build payer spend breakdown for response
	var payerPointsToSpend = calculatePayerPointsToSpend(balancesAscending, pointsToSpend)

	// add transactions to transactions array
	AddTransactions(payerPointsToSpend)

	// return 201 response
	context.JSON(http.StatusAccepted, payerPointsToSpend)
}

// consolidateTransactions takes transactions and reduces to positive balances then returns
func consolidateTransactions() []models.Transaction {
	var balancesAscending []models.Transaction
	transactions := getTransactions()

	// loop transactions and get actual balances after negatives
	for _, transaction := range transactions {

		// checks for positive transactions and appends to balancesAscending then moves on
		if transaction.Points > 0 {
			balancesAscending = append(balancesAscending, transaction)
			continue
		}

		balancesAscending = balancesAscendingDeduct(balancesAscending, transaction)
	}

	return balancesAscending
}

// reduces transaction points from balanceAscending records based on matching payers, returns updated array
func balancesAscendingDeduct(balancesAscending []models.Transaction, transaction models.Transaction) []models.Transaction {
	// loop through balance ascending and subtract from payer balances oldest to newest
	for i := 0; i < len(balancesAscending); i++ {
		balance := balancesAscending[i]

		// skips non-matching payers and balance points less than 0
		if balance.Payer != transaction.Payer || balance.Points <= 0 {
			continue
		}

		// convert to positive since it comes in as negative
		transactionPoints := transaction.Points * -1

		// if balance points satisfy transaction points deduct from balance and break
		if balance.Points >= transactionPoints {
			// handles balances that satisfy entire transaction
			balancesAscending[i].Points = balance.Points - transactionPoints
			break
		}

		// if transaction is more than points update transaction points and unset balance
		transaction.Points = transactionPoints - balance.Points
		balancesAscending[i].Points = 0
	}

	return balancesAscending
}

// takes positive balance array and points to reduce, goes through and breaks out points to be spent by payer
func calculatePayerPointsToSpend(balancesAscending []models.Transaction, pointsToSpend models.SpendPointsInput) []models.PayerSpendPoints {
	// initialize response
	var payerPointsToSpend []models.PayerSpendPoints

	// go through balances and build payer => points map
	for i := 0; i < len(balancesAscending); i++ {
		balance := balancesAscending[i]

		// skip 0 balances
		if balance.Points == 0 {
			continue
		}

		// Track points to add to payer points
		pointsToAdd := 0

		// calculate how many points will be spent from balance and if all balance points are used
		if balance.Points > pointsToSpend.Points {
			pointsToAdd = pointsToSpend.Points
		} else {
			pointsToAdd = balance.Points
			balancesAscending[i].Points = 0
		}

		// gets payer index, if not found will add
		payerIndex := getOrAddPayerIndex(&payerPointsToSpend, balance.Payer)

		// add points to payer and remove from points needed
		payerPointsToSpend[payerIndex].Points += pointsToAdd * -1
		pointsToSpend.Points += pointsToAdd * -1

		// if points needed have been met, break out of loop
		if pointsToSpend.Points == 0 {
			break
		}
	}

	return payerPointsToSpend
}

// payerExists gets or adds payer index in array of payers
func getOrAddPayerIndex(payersArray *[]models.PayerSpendPoints, value string) int {
	// returns payers index if found
	for index, payer := range *payersArray {
		if payer.Payer == value {
			return index
		}
	}

	// creates payer if not found
	*payersArray = append(*payersArray, models.PayerSpendPoints{Payer: value, Points: 0})

	// return added payer index
	return len(*payersArray) - 1
}
