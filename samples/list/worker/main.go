package main

import (
	"context"
	"log"

	"go.temporal.io/sdk/worker"

	sampleworkflow "github.com/Zampfi/workflow/sdk/samples/list/sampleworkflow"
	"github.com/Zampfi/workflow/sdk/workflowmanagers/temporal"
	"github.com/Zampfi/workflow/sdk/workflowmanagers/temporal/models"
	"github.com/Zampfi/workflow/sdk/workflowmanagers/temporal/workflow"
)

func main() {
	ctx := context.Background()
	temporalService := temporal.NewTemporalService()
	err := temporalService.Connect(ctx, models.ConnectClientParams{})
	if err != nil {
		panic(err)
	}

	w, err := temporalService.GetNewWorker(ctx, models.NewWorkerParams{
		TaskQueue: "test_execute",
		Workflows: []workflow.Workflow{
			{
				Function: sampleworkflow.SampleWorkflow,
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
