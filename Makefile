# Makefile

.PHONY: infra.start workflow.test  

# Start the infrastructure
infra.start:
	docker-compose up --build -d

# Run tests
workflow.test:
	ginkgo test -environment=qa
