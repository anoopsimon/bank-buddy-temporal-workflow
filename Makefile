# Makefile

.PHONY: help infra.start workflow.test  

# Start the infrastructure
infra.start:
	docker-compose up --build -d

# Run tests
workflow.test:
	ginkgo test -environment=qa

.PHONY: help

help:
	@echo Available commands:
	@echo infra.start     - Start the infrastructure
	@echo workflow.test   - Run the workflow tests
