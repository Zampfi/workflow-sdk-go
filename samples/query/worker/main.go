package main

import (
	"context"
	"log"

	querysampleworkflow "github.com/Zampfi/workflow-sdk-go/sdk/samples/query/sampleworkflow"
	"github.com/Zampfi/workflow-sdk-go/sdk/workflowmanagers/temporal"
	"github.com/Zampfi/workflow-sdk-go/sdk/workflowmanagers/temporal/models"
	"github.com/Zampfi/workflow-sdk-go/sdk/workflowmanagers/temporal/workflow"
	"go.temporal.io/sdk/worker"
)

func main() {
	ctx := context.Background()
	temporalService := temporal.NewTemporalService()
	err := temporalService.Connect(ctx, models.ConnectClientParams{})
	if err != nil {
		panic(err)
	}

	w, err := temporalService.GetNewWorker(ctx, models.NewWorkerParams{
		TaskQueue: "test_query",
		Workflows: []workflow.Workflow{
			{
				Function: querysampleworkflow.SampleWorkflow,
			},
		},
		RegisterTasks: true,
		Options: worker.Options{
			WorkerActivitiesPerSecond: 10,
		},
	})

	if err != nil {
		log.Fatalln("Unable to create worker", err)
	}

	err = w.Run(ctx, worker.InterruptCh(), models.RunWorkerParams{})
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
