package main

import (
	"context"
	"fmt"
)

func SendAlertActivity(ctx context.Context, paymentID string) error {
	fmt.Println("[ALERT] Pembayaran terlambat:", paymentID)
	return nil
}
