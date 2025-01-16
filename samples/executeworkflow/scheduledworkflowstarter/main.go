package main

import (
	"context"
	"fmt"

	"github.com/Zampfi/workflow-sdk-go/workflowmanagers/temporal"
	"github.com/Zampfi/workflow-sdk-go/workflowmanagers/temporal/models"
	"go.temporal.io/sdk/client"
)

func main() {
	ctx := context.Background()
	temporalService := temporal.NewTemporalService()
	err := temporalService.Connect(ctx, models.ConnectClientParams{})
	if err != nil {
		panic(err)
	}

	asyncRun, err := temporalService.ExecuteScheduledWorkflow(
		ctx,
		models.ExecuteWorkflowWithScheduleParams{
			ScheduleOptions: client.ScheduleOptions{
				ID: "ScheduleTestNow",
				Spec: client.ScheduleSpec{
					CronExpressions: []string{"* * * * *"},
				},
				Action: &client.ScheduleWorkflowAction{
					ID:        "test-workflow",
					Workflow:  "SampleWorkflow",
					Args:      []interface{}{2},
					TaskQueue: "test_execute",
				},
				PauseOnFailure:     true,
				TriggerImmediately: true,
			},
		},
	)

	if err != nil {
		fmt.Printf("Error executing workflow: %v\n", err)
	}
	fmt.Printf("Async run: %v\n", asyncRun)

}
