package test

import (
	"context"
	"log"
	"os"
	"path/filepath"

	"github.com/yourusername/temporal-sample-workflow/workflows"
	"go.temporal.io/sdk/client"
	"gopkg.in/yaml.v2"
)

var ENVIRONMENT = "qa"

// QA is the namespace for testing helpers.
var QA = struct {
	Temporal struct {
		TriggerCustomerCreationWorkflow func(input CustomerCreationInput) (string, error)
	}
}{}

type CustomerCreationInput struct {
	CustomerName  string
	CustomerID    int
	DepositAmount float64
}

func init() {
	config, err := LoadConfig(ENVIRONMENT)
	if err != nil {
		log.Fatalln("Unable to load configuration:", err)
	}

	QA.Temporal.TriggerCustomerCreationWorkflow = func(input CustomerCreationInput) (string, error) {
		c, err := client.NewClient(client.Options{})
		if err != nil {
			log.Fatalln("Unable to create Temporal client", err)
		}
		defer c.Close()

		options := client.StartWorkflowOptions{
			ID:        config.Temporal.ID,        // Use the ID from the config
			TaskQueue: config.Temporal.TaskQueue, // Use the Task Queue from the config
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

// Config holds the configuration values from the YAML file.
type Config struct {
	Temporal struct {
		ID        string `yaml:"id"`
		TaskQueue string `yaml:"task_queue"`
	} `yaml:"temporal"`
}

// LoadConfig reads the YAML configuration file for the specified environment.
func LoadConfig(env string) (*Config, error) {
	// Get the project root directory by navigating upwards
	rootDir, err := findProjectRoot()
	if err != nil {
		return nil, err
	}

	// Construct the config file path
	configFile := filepath.Join(rootDir, "config", env+".yaml")
	file, err := os.Open(configFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

// findProjectRoot traverses up the directory tree to find the project root.
func findProjectRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if isProjectRoot(dir) {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir { // Reached the root of the filesystem
			break
		}
		dir = parent
	}

	return "", os.ErrNotExist // Project root not found
}

// isProjectRoot checks if the current directory is the project root.
func isProjectRoot(dir string) bool {
	// Check for a specific file or directory that indicates project root, e.g., go.mod
	if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
		return true
	}
	// You can add more checks here if needed (like for "config" folder)
	return false
}
