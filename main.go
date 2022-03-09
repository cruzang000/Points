package main

import (
	transactions2 "Points/api/controllers/transactions"
	"github.com/gin-gonic/gin"
)

func main() {
	router := SetupRouter()

	err := router.Run("localhost:8080")

	if err != nil {
		return
	}
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/addTransaction", transactions2.AddTransaction)
	router.PUT("/spendPoints", transactions2.SpendPoints)
	router.GET("/getBalances", transactions2.GetBalances)

	// Handle error response when a route is not defined
	router.NoRoute(func(c *gin.Context) {
		// In gin this is how you return a JSON response
		c.JSON(404, gin.H{"message": "Not found"})
	})

	return router
}
