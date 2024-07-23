package domain

import "time"

type Payment struct {
	ID         int64       
	CreatedAt  int64
	UserId     int32
	OrderId    int32
	TotalPrice float32
}

func NewPayment(userId, orderId int32, totalPrice float32) *Payment {
	return &Payment{
		CreatedAt:  time.Now().Unix(),
		UserId:     userId,
		OrderId:    orderId,
		TotalPrice: totalPrice,
	}
}
