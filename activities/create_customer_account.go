package activities

import "fmt"

// CreateCustomerAccountActivity simulates creating a new account for the customer.
func CreateCustomerAccountActivity(customerID int) (string, error) {
	accountInfo := fmt.Sprintf("Account created for Customer ID: %d with account number: 12345678", customerID)
	return accountInfo, nil
}
