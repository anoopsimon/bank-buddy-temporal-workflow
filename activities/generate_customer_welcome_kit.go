package activities

import "fmt"

// GenerateCustomerWelcomeKitActivity generates a welcome kit for the new customer.
func GenerateCustomerWelcomeKitActivity(customerName string) (string, error) {
	message := fmt.Sprintf("Welcome kit generated for customer: %s.", customerName)
	return message, nil
}
