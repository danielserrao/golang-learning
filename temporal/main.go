package main

import (
	"context"
	"errors"
	"fmt"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/workflow"
)

// Define the workflow function
func SimpleWorkflow(ctx workflow.Context) error {
	logger := workflow.GetLogger(ctx)
	logger.Info("SimpleWorkflow started")

	return errors.New("testing workflow fail")
	//// Here you can define your workflow logic
	//workflow.Sleep(ctx, time.Second*10) // Sleep for 10 seconds
	//
	//logger.Info("SimpleWorkflow completed")
	//return nil
}

func main() {
	// Create a Temporal client
	c, err := client.NewLazyClient(client.Options{
		HostPort:  "localhost:7233",
		Namespace: "cortex-portal",
	})

	if err != nil {
		fmt.Println("Unable to create Temporal client", err)
		return
	}
	defer c.Close()

	// Start a workflow execution
	workflowOptions := client.StartWorkflowOptions{
		ID:        "simple_workflow_test3",
		TaskQueue: "v5.47.1-wrong4",
	}

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, SimpleWorkflow)
	if err != nil {
		fmt.Println("Unable to execute workflow", err)
		return
	}

	// Print the workflow execution details
	fmt.Printf("Started workflow with WorkflowID: %s and RunID: %s\n", we.GetID(), we.GetRunID())
}
