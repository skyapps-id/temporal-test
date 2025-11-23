package workflow

import "fmt"

func SendPaymentReminderActivity(orderID string) error {
	fmt.Println("Push notif: ", orderID)
	// TODO push notif

	return nil
}

func CheckPaymentStatusActivity(orderID string) (bool, error) {
	// TODO Query DB or API

	// return true if paid, false if not
	return false, nil
}

func CancelOrderActivity(orderID string) error {
	// TODO update DB: order canceled and push notif

	return nil
}
