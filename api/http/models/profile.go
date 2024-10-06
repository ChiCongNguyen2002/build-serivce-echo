package models

import "time"

type GetUserTransactionHistoryReq struct {
	ProfileID   string `query:"profileID"`
	Offset      int64  `query:"offset"`
	Limit       int64  `query:"limit"`
	TxType      string `query:"txType"`
	Status      string `query:"status"`
	RecentMonth int    `query:"recentMonth"`
}

type GetUserTransactionHistoryByProfileReq struct {
	ProfileID string `query:"profileID"`
}

type UserTransactionHistory struct {
	TransactionID        string     `json:"transactionID"`
	TransactionType      string     `json:"transactionType"`
	ProfileID            string     `json:"profileID"`
	Status               string     `json:"status"`
	PointAmount          int64      `json:"pointAmount"`
	PointType            int64      `json:"pointType"`
	TotalAmount          float64    `json:"totalAmount"`
	Currency             string     `json:"currency"`
	PaymentTransactionID string     `json:"paymentTransactionID"`
	Source               string     `json:"source"`
	SourceTime           *time.Time `json:"sourceTime"`
	SourceType           string     `json:"sourceType"`
	CreatedAt            *time.Time `json:"createdAt"`
	UpdatedAt            *time.Time `json:"updatedAt"`
}
