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

	asyncRun, err := temporalService.ExecuteAsyncWorkflow(ctx, models.ExecuteWorkflowParams{
		Workflow: "SampleWorkflow",
		Args:     []interface{}{2},
		Options: models.StartWorkflowOptions{
			TaskQueue: "test_execute",
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

}
