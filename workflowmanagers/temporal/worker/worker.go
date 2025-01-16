package worker

import (
	"context"
	"errors"

	activity "github.com/Zampfi/workflow-sdk-go/sdk/workflowmanagers/temporal/activity"
	"github.com/Zampfi/workflow-sdk-go/sdk/workflowmanagers/temporal/client"
	"github.com/Zampfi/workflow-sdk-go/sdk/workflowmanagers/temporal/models"
	workflow "github.com/Zampfi/workflow-sdk-go/sdk/workflowmanagers/temporal/workflow"
	temporalworker "go.temporal.io/sdk/worker"
)

type Worker interface {
	Run(ctx context.Context, interruptCh <-chan interface{}, params models.RunWorkerParams) error
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}

type worker struct {
	temporalworker  temporalworker.Worker
	taskQueue       string
	client          *client.TemporalClient
	isRegisterTasks bool
	workflows       []workflow.Workflow
	activities      []activity.Activity
}

func NewWorker(ctx context.Context, tc *client.TemporalClient, params models.NewWorkerParams) (Worker, error) {
	if tc == nil {
		return nil, errors.New("temporal client is required")
	}

	baseClient := tc.GetBaseClient()
	if baseClient == nil {
		return nil, errors.New("base Temporal client is required")
	}

	w := temporalworker.New(baseClient, params.TaskQueue, params.ToTemporalWorkerOptions())

	return &worker{
		temporalworker:  w,
		taskQueue:       params.TaskQueue,
		client:          tc,
		isRegisterTasks: params.RegisterTasks,
		workflows:       params.Workflows,
		activities:      params.Activities,
	}, nil
}

func (w *worker) Run(ctx context.Context, interruptCh <-chan interface{}, params models.RunWorkerParams) error {
	if w.isRegisterTasks {
		w.registerTasks()
	}
	return w.temporalworker.Run(interruptCh)
}

func (w *worker) registerTasks() {
	for _, workflow := range w.workflows {
		w.temporalworker.RegisterWorkflowWithOptions(workflow.Function, workflow.RegisterOptions)
	}

	for _, activity := range w.activities {
		w.temporalworker.RegisterActivityWithOptions(activity.Function, activity.RegisterOptions)
	}
}

func (w *worker) Start(ctx context.Context) error {
	return w.temporalworker.Start()
}

func (w *worker) Stop(ctx context.Context) error {
	w.temporalworker.Stop()
	return nil
}
