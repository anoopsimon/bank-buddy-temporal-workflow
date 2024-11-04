package main

import (
	"context"
	"log"

	"github.com/yourusername/temporal-sample-workflow/workflows"
	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "customer-creation-workflow",
		TaskQueue: "customer-task-queue",
	}

	we, err := c.ExecuteWorkflow(context.Background(), options, workflows.CustomerCreationWorkflow, "John Doe", 101, 500.0)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}

	log.Printf("Started workflow with WorkflowID: %s and RunID: %s", we.GetID(), we.GetRunID())

	var result string
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable to get workflow result", err)
	}

	log.Printf("Workflow result: %s", result)
}
