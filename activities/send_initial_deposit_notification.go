package activities

import "fmt"

// SendInitialDepositNotificationActivity notifies the customer of their initial deposit.
func SendInitialDepositNotificationActivity(customerName string, depositAmount float64) (string, error) {
	message := fmt.Sprintf("Hello %s, an initial deposit of $%.2f has been added to your account.", customerName, depositAmount)
	return message, nil
}
