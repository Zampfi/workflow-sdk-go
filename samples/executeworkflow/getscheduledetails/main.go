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

	querySchedule, err := temporalService.QuerySchedule(ctx, models.QueryScheduleParams{
		ScheduleID: "ScheduleID",
	})
	if err != nil {
		panic(err)
	}

	workflowID := querySchedule.ScheduleDescription.Schedule.Action.(*client.ScheduleWorkflowAction).ID

	fmt.Printf("Query schedule: %v\n", workflowID)

	result, err := temporalService.ListWorkflows(ctx, models.ListWorkflowsParams{
		PageSize: 1000,
		Query:    fmt.Sprintf("TemporalScheduledById = '%s'", "ScheduleID"),
	})

	println(result.Workflows)

}
