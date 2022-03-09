package main

import (
	"Points/api/models"
	"Points/utils"
	"github.com/gin-gonic/gin"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var router *gin.Engine
var testTransactions []models.Transaction

// init runs before test methods to populate transactions
func init() {
	router = SetupRouter()

	// get test transactions
	addTestTransactions()

	// populate some balances using transactions
	for index := range testTransactions {
		transaction := testTransactions[index]
		payload := models.Transaction{Payer: transaction.Payer, Points: transaction.Points, Timestamp: transaction.Timestamp}

		rr, req := utils.HttpTestRequest("POST", "/addTransaction", utils.Stringify(payload))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(rr, req)
	}
}

// TestGetBalances tests addBalance api endpoint
func TestGetBalances(t *testing.T) {
	// send test api request for balance
	rr, req := utils.HttpTestRequest("GET", "/getBalances", nil)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(rr, req)

	// parse balance string
	response := string(rr.Body.Bytes())

	// compare expected string to response string
	assert.Equal(t, "{\"DANNON\":1100,\"MILLER COORS\":10000,\"UNILEVER\":200}", response)
}

// TestAddBalance tests addBalance api endpoint
func TestAddBalance(t *testing.T) {
	// create test transaction
	tm4, _ := time.Parse(time.RFC3339, "2022-03-02T14:00:00Z")
	singleTestTransaction := models.Transaction{Payer: "DANNON", Points: 5000, Timestamp: tm4}

	// send test api request
	rr, req := utils.HttpTestRequest("POST", "/addTransaction", utils.Stringify(singleTestTransaction))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(rr, req)

	// process response as string
	response := string(rr.Body.Bytes())

	// compare expected string to response string
	assert.Equal(t, "{\"payer\":\"DANNON\",\"points\":5000,\"timestamp\":\"2022-03-02T14:00:00Z\"}", response)
}

// TestSpendPoints tests points points
func TestSpendPoints(t *testing.T) {
	spendPointsInput := models.SpendPointsInput{Points: 5000}

	// send test api request
	rr, req := utils.HttpTestRequest("PUT", "/spendPoints", utils.Stringify(spendPointsInput))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(rr, req)

	// process response as string
	response := string(rr.Body.Bytes())

	// compare expected string to response string
	assert.Equal(t, "[{\"Payer\":\"DANNON\",\"Points\":-100},{\"Payer\":\"UNILEVER\",\"Points\":-200},{\"Payer\":\"MILLER COORS\",\"Points\":-4700}]", response)
}

// addTestTransactions adds test transactions to testTransactions slice
func addTestTransactions() {
	testTransactions = nil

	tm, _ := time.Parse(time.RFC3339, "2020-10-31T10:00:00Z")
	tm1, _ := time.Parse(time.RFC3339, "2020-10-31T11:00:00Z")
	tm2, _ := time.Parse(time.RFC3339, "2020-10-31T15:00:00Z")
	tm3, _ := time.Parse(time.RFC3339, "2020-11-01T14:00:00Z")
	tm4, _ := time.Parse(time.RFC3339, "2020-11-02T14:00:00Z")

	testTransactions = append(testTransactions, models.Transaction{Payer: "DANNON", Points: 1000, Timestamp: tm4})
	testTransactions = append(testTransactions, models.Transaction{Payer: "UNILEVER", Points: 200, Timestamp: tm1})
	testTransactions = append(testTransactions, models.Transaction{Payer: "MILLER COORS", Points: 10000, Timestamp: tm3})
	testTransactions = append(testTransactions, models.Transaction{Payer: "DANNON", Points: -200, Timestamp: tm2})
	testTransactions = append(testTransactions, models.Transaction{Payer: "DANNON", Points: 300, Timestamp: tm})
}
