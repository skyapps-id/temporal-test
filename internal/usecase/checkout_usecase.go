package usecase

import (
	"temporal-test/internal/starter"
)

type CheckoutUsecase struct {
	Starter *starter.TemporalStarter
}

func NewCheckoutUsecase(st *starter.TemporalStarter) *CheckoutUsecase {
	return &CheckoutUsecase{Starter: st}
}

func (u *CheckoutUsecase) Checkout(orderID string, amount int64) error {
	return u.Starter.StartPaymentWorkflow(orderID, amount)
}
