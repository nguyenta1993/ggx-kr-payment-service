package entity

type Transaction struct {
	ID            int64  `json:"id"`
	OrderId       int64  `json:"orderId"`
	TransactionID string `json:"transactionId"`
	OrderSnapShot string `json:"orderSnapShot"`
	CreatedAt     int64  `json:"createdAt"`
	UpdatedAt     int64  `json:"updatedAt"`
	Status        string `json:"status"`
}
