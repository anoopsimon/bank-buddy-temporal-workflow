# Makefile

.PHONY: help infra.start workflow.test stack.lint  

# Start the infrastructure
infra.start:
	docker-compose up --build -d 

# Run tests
workflow.test:
	ginkgo --json-report=test-results.json test -environment=qa 

# Linting command
stack.lint:
	docker-compose exec lint golangci-lint run

.PHONY: help

help:
	@echo Available commands:
	@echo infra.start     - Start the infrastructure
	@echo workflow.test   - Run the workflow tests
	@echo stack.lint      - Run the linter
