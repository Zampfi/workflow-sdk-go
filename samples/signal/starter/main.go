package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Zampfi/workflow-sdk-go/sdk/workflowmanagers/temporal"
	"github.com/Zampfi/workflow-sdk-go/sdk/workflowmanagers/temporal/models"
	"github.com/lithammer/shortuuid"
	temporalmain "go.temporal.io/sdk/temporal"
)

func main() {
	ctx := context.Background()
	temporalService := temporal.NewTemporalService()
	err := temporalService.Connect(ctx, models.ConnectClientParams{})
	if err != nil {
		panic(err)
	}

	// Start a new workflow execution
	workflowID := fmt.Sprintf("SampleSignalWorkflow-%s", shortuuid.New())
	run, err := temporalService.ExecuteAsyncWorkflow(ctx, models.ExecuteWorkflowParams{
		Workflow: "SampleSignalWorkflow",
		Options: models.StartWorkflowOptions{
			TaskQueue: "signal-queue",
			ID:        workflowID,
			RetryPolicy: &temporalmain.RetryPolicy{
				InitialInterval:    1 * time.Second,
				BackoffCoefficient: 2,
				MaximumInterval:    1 * time.Minute,
				MaximumAttempts:    5,
			},
		},
	})
	if err != nil {
		fmt.Printf("Error executing workflow: %v\n", err)
	}
	fmt.Printf("Workflow started: %v\n", run.WorkflowID)
}
