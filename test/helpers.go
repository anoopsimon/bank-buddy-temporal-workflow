package test

import (
    "context"
    "log"
    "go.temporal.io/sdk/client"
    "github.com/yourusername/temporal-sample-workflow/workflows"
)

// QA is the namespace for testing helpers.
var QA = struct {
    Temporal struct {
        TriggerCustomerCreationWorkflow func(input CustomerCreationInput) (string, error)
    }
}{}

// CustomerCreationInput holds the parameters to trigger the workflow.
type CustomerCreationInput struct {
    CustomerName   string
    CustomerID     int
    DepositAmount  float64
}

// init initializes the QA helpers for triggering workflows.
func init() {
    QA.Temporal.TriggerCustomerCreationWorkflow = func(input CustomerCreationInput) (string, error) {
        c, err := client.NewClient(client.Options{})
        if err != nil {
            log.Fatalln("Unable to create Temporal client", err)
        }
        defer c.Close()

        options := client.StartWorkflowOptions{
            ID:        "customer-creation-workflow-test",
            TaskQueue: "customer-task-queue",
        }

        we, err := c.ExecuteWorkflow(context.Background(), options, workflows.CustomerCreationWorkflow, input.CustomerName, input.CustomerID, input.DepositAmount)
        if err != nil {
            return "", err
        }

        var result string
        err = we.Get(context.Background(), &result)
        return result, err
    }
}
