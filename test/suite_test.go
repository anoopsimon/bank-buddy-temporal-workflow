package test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

func TestCustomerCreationWorkflow(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CustomerCreationWorkflow Suite")
}
