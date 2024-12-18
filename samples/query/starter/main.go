package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Zampfi/citadel/workflowmanagers/temporal"
	"github.com/Zampfi/citadel/workflowmanagers/temporal/models"
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

	syncRun, err := temporalService.ExecuteSyncWorkflow(ctx, models.ExecuteWorkflowParams{
		Workflow: "SampleWorkflow",
		Options: models.StartWorkflowOptions{
			TaskQueue: "test_query",
			ID:        fmt.Sprintf("SampleWorkflow-%s", shortuuid.New()),
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
	fmt.Printf("Sync run: %v\n", syncRun)

	asyncRun, err := temporalService.ExecuteAsyncWorkflow(ctx, models.ExecuteWorkflowParams{
		Workflow: "SampleWorkflow",
		Options: models.StartWorkflowOptions{
			TaskQueue: "test_query",
			ID:        fmt.Sprintf("SampleWorkflow-%s", shortuuid.New()),
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
	fmt.Printf("Async run: %v\n", asyncRun)

	var syncCounter int
	querySyncWorkflow, err := temporalService.QueryWorkflow(ctx, models.QueryWorkflowParams{
		WorkflowID: syncRun.WorkflowID,
		QueryType:  "getCounter",
	})
	if err != nil {
		fmt.Printf("Error querying workflow: %v\n", err)
	}

	querySyncWorkflow.QueryResult.Get(&syncCounter)
	fmt.Printf("Query response: %v\n", querySyncWorkflow)
	fmt.Println()
	fmt.Printf("SyncCounter: %v\n", syncCounter)

	var asyncCounter int
	queryAsyncWorkflow, err := temporalService.QueryWorkflow(ctx, models.QueryWorkflowParams{
		WorkflowID: asyncRun.WorkflowID,
		QueryType:  "getCounter",
	})

	if err != nil {
		fmt.Printf("Error querying workflow: %v\n", err)
	}

	queryAsyncWorkflow.QueryResult.Get(&asyncCounter)
	fmt.Printf("Query response: %v\n", queryAsyncWorkflow)
	fmt.Println()
	fmt.Printf("AsyncCounter: %v\n", asyncCounter)
}
