package main

import (
	"log"

	"github.com/yourusername/temporal-sample-workflow/activities"
	"github.com/yourusername/temporal-sample-workflow/workflows"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.NewClient(client.Options{
		HostPort: "temporal:7233",
	})
	if err != nil {
		log.Fatalln("Unable to create Temporal client", err)
	}
	defer c.Close()

	w := worker.New(c, "customer-task-queue", worker.Options{})
	w.RegisterWorkflow(workflows.CustomerCreationWorkflow)
	w.RegisterActivity(activities.SendWelcomeEmailActivity)
	w.RegisterActivity(activities.CreateCustomerAccountActivity)
	w.RegisterActivity(activities.SendInitialDepositNotificationActivity)
	w.RegisterActivity(activities.GenerateCustomerWelcomeKitActivity)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
