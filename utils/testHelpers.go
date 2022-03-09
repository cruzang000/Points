package utils

import (
	"Points/api/models"
	"encoding/json"
	"github.com/sirupsen/logrus"
)

type PayerPointsStruct struct {
	Payer  string `json:"payer" binding:"required" `
	Points int    `json:"points" binding:"required"`
}

func Stringify(payload interface{}) []byte {
	response, _ := json.Marshal(payload)
	return response
}

func ParseSingleTransaction(payload []byte) models.Transaction {
	var singleTransactionResponse models.Transaction

	err := json.Unmarshal(payload, &singleTransactionResponse)

	if err != nil {
		logrus.Error(err.Error())
	}

	return singleTransactionResponse
}

func ParseSpendPointsResponse(payload []byte) PayerPointsStruct {
	var spendPointsResponse PayerPointsStruct

	err := json.Unmarshal(payload, &spendPointsResponse)

	if err != nil {
		logrus.Error(err.Error())
	}

	return spendPointsResponse
}
