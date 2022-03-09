package models

import "time"

type Transaction struct {
	Payer     string    `json:"payer" binding:"required" `
	Points    int       `json:"points" binding:"required"`
	Timestamp time.Time `json:"timestamp" binding:"required"`
}

type SpendPointsInput struct {
	Points int `json:"points" binding:"required"`
}

type PayerSpendPoints struct {
	Payer  string
	Points int
}
