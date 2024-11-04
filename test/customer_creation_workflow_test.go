package test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("CustomerCreationWorkflow", func() {
	It("should complete successfully and return the correct result", func() {
		input := CustomerCreationInput{
			CustomerName:  "John Doe",
			CustomerID:    101,
			DepositAmount: 500.0,
		}

		result, err := QA.Temporal.TriggerCustomerCreationWorkflow(input)
		Expect(err).To(BeNil())
		Expect(result).To(ContainSubstring("Workflow completed successfully"))
	})
})
