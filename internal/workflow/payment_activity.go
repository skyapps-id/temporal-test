package workflow

import "fmt"

func SendPaymentReminderActivity(orderID string) error {
	fmt.Println("Push notif reminder: ", orderID)

	// TODO push notif

	return nil
}

func CheckPaymentStatusActivity(orderID string) (bool, error) {
	fmt.Println("Check payment status: ", orderID)

	// TODO Query DB or API

	// return true if paid, false if not
	return false, nil
}

func CancelOrderActivity(orderID string) error {
	fmt.Println("Push notif cancel order : ", orderID)

	// TODO update DB: order canceled and push notif

	return nil
}
