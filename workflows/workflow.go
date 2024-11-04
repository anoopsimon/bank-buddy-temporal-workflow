package workflows

import (
	"fmt"
	"time"

	"github.com/yourusername/temporal-sample-workflow/activities"
	"go.temporal.io/sdk/workflow"
)

// CustomerCreationWorkflow simulates a workflow that performs post-customer creation tasks.
func CustomerCreationWorkflow(ctx workflow.Context, customerName string, customerID int, depositAmount float64) (string, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 10,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	var results []string

	var welcomeMsg string
	if err := workflow.ExecuteActivity(ctx, activities.SendWelcomeEmailActivity, customerName).Get(ctx, &welcomeMsg); err != nil {
		return "", err
	}
	results = append(results, welcomeMsg)

	var accountInfo string
	if err := workflow.ExecuteActivity(ctx, activities.CreateCustomerAccountActivity, customerID).Get(ctx, &accountInfo); err != nil {
		return "", err
	}
	results = append(results, accountInfo)

	var depositNotification string
	if err := workflow.ExecuteActivity(ctx, activities.SendInitialDepositNotificationActivity, customerName, depositAmount).Get(ctx, &depositNotification); err != nil {
		return "", err
	}
	results = append(results, depositNotification)

	var welcomeKit string
	if err := workflow.ExecuteActivity(ctx, activities.GenerateCustomerWelcomeKitActivity, customerName).Get(ctx, &welcomeKit); err != nil {
		return "", err
	}
	results = append(results, welcomeKit)

	// Combine all results for the final output
	return fmt.Sprintf("Workflow completed successfully with results: %v", results), nil
}
