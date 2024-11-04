package activities

import "fmt"

// SendWelcomeEmailActivity simulates sending a welcome email after customer creation.
func SendWelcomeEmailActivity(customerName string) (string, error) {
	message := fmt.Sprintf("Welcome to the bank, %s! We're excited to have you as a customer.", customerName)
	return message, nil
}
