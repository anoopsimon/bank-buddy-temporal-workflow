# Temporal


# Temporal Customer Creation Workflow

This project demonstrates a simple Temporal workflow to simulate customer creation tasks in a bank. The workflow includes multiple activities that represent common post-creation processes, such as account initialization and welcome package delivery.


## Tech Stack

| Tool                                                                 | Description                                        |
|----------------------------------------------------------------------|----------------------------------------------------|
| <img alt="Go" src="https://img.shields.io/badge/-Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" /> | Programming Language                               |
| <img alt="Docker" src="https://img.shields.io/badge/-Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white" /> | Containerization Platform                           |
| <img alt="Docker Compose" src="https://img.shields.io/badge/-Docker%20Compose-2496ED?style=for-the-badge&logo=docker&logoColor=white" /> | Tool for defining and running multi-container applications |
| <img alt="Temporal" src="https://img.shields.io/badge/-Temporal-00B3B3?style=for-the-badge&logo=temporal&logoColor=white" /> | Workflow Orchestration Tool                        |
| <img alt="Ginkgo" src="https://img.shields.io/badge/-Ginkgo-00C4A1?style=for-the-badge&logo=ginkgo&logoColor=white" /> | BDD Testing Framework                              |
| <img alt="Gomega" src="https://img.shields.io/badge/-Gomega-00C4A1?style=for-the-badge&logo=gomega&logoColor=white" /> | Matcher library for Go tests                       |


## Prerequisites

Before running the project, make sure you have the following installed:

- **Go** (version 1.17 or higher): [Download Go](https://golang.org/dl/)
- **Docker**: [Download Docker](https://www.docker.com/get-started)
- **Docker Compose**: Typically included with Docker Desktop. Verify installation by running `docker-compose --version`.
- **Make**: Optional: If you want to run infra and tests without using docker commands . Install from internet and run command to verify. `make --version`.

## Setting Up Temporal with Docker Compose

To spin up a Temporal server and UI using Docker Compose, use the following setup:

1. Start Temporal services with:

   ```bash
   docker-compose up --build -d
   # Alternatively use make command
   make infra.start
   ```
Following services will be running in inside the container
```js
[+] Running 6/6
 ✔ Container temporal-elasticsearch  Running                                                                                                                    0.0s 
 ✔ Container temporal-postgresql     Running                                                                                                                    0.0s 
 ✔ Container temporal                Running                                                                                                                    3.8s 
 ✔ Container temporal-worker         Running                                                                                                                    0.4s 
 ✔ Container temporal-admin-tools    Running                                                                                                                    1.5s 
 ✔ Container temporal-ui             Running 
   ```

3. Access the Temporal UI by navigating to [Temporal UI](http://localhost:8080) in your web browser.

## Scope of the Temporal Workflow

The `CustomerCreationWorkflow` workflow performs the following activities as part of the customer creation process:

1. **InitializeAccountActivity**: Initializes the customer's account with basic information.
2. **SendWelcomePackageActivity**: Sends a welcome package to the new customer.
3. **SetUpDirectDepositActivity**: Configures direct deposit for the customer.
4. **ActivateDebitCardActivity**: Activates a debit card linked to the customer's account.

These activities are managed in a sequence within the Temporal workflow.

## Test Suite

The project includes a Ginkgo test suite to validate the workflow’s functionality. The test suite is located in the `test/` directory.

### Installing the Ginkgo Tool

To install the Ginkgo CLI tool, run:

```bash
go install github.com/onsi/ginkgo/v2/ginkgo@latest
```

### Running Tests

To run the Ginkgo tests, use the following command from the project root:

```bash
ginkgo --json-report=test-results.json test -environment=qa
# or if you are using make command
make test
```

Sample output:

```plaintext
PS C:\Users\s4514\dev\go\bank-buddy-temporal-workflowe> ginkgo -r test/
Running Suite: CustomerCreationWorkflow Suite - C:\Users\s4514\dev\go\bank-buddy-temporal-workflow\test
=========================================================================================
Random Seed: 1730705962

Will run 1 of 1 specs
2024/11/04 18:40:36 INFO  No logger configured for temporal client. Created default one.
+

Ran 1 of 1 Specs in 0.811 seconds
SUCCESS! -- 1 Passed | 0 Failed | 0 Pending | 0 Skipped
PASS

Ginkgo ran 1 suite in 1m14.8897583s
Test Suite Passed
```

This verifies that the `CustomerCreationWorkflow` executes successfully and completes all activities as expected.

### Temporal workflow screenshot
![Temporal Workflow](doc\img\workflow_scr.png)


### Containers
![Temporal Workflow](doc\img\containers.png)

