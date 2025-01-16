package temporal

import (
	"context"
	"sync"

	"github.com/Zampfi/workflow-sdk-go/sdk/workflowmanagers/temporal/client"
	"github.com/Zampfi/workflow-sdk-go/sdk/workflowmanagers/temporal/models"
	worker "github.com/Zampfi/workflow-sdk-go/sdk/workflowmanagers/temporal/worker"
)

type TemporalService interface {
	Connect(ctx context.Context, params models.ConnectClientParams) error
	ExecuteAsyncWorkflow(ctx context.Context, params models.ExecuteWorkflowParams) (models.WorkflowResponse, error)
	ExecuteSyncWorkflow(ctx context.Context, params models.ExecuteWorkflowParams) (models.WorkflowResponse, error)
	ListWorkflows(ctx context.Context, params models.ListWorkflowsParams) (models.ListWorkflowsResponse, error)
	GetWorkflowDetails(ctx context.Context, params models.GetWorkflowDetailsParams) (models.WorkflowDetailsResponse, error)
	TerminateWorkflow(ctx context.Context, params models.TerminateWorkflowParams) error
	CancelWorkflow(ctx context.Context, params models.CancelWorkflowParams) error
	SignalWorkflow(ctx context.Context, params models.SignalWorkflowParams) error
	QueryWorkflow(ctx context.Context, params models.QueryWorkflowParams) (models.QueryWorkflowResponse, error)
	GetNewWorker(ctx context.Context, params models.NewWorkerParams) (worker.Worker, error)
	Close(ctx context.Context)
}

type temporalService struct {
	client    client.Client
	connected bool
	mu        sync.Mutex
}

func NewTemporalService() TemporalService {
	return &temporalService{
		client: &client.TemporalClient{},
	}
}

// This is to test mock client for testing only.
func (ts *temporalService) setClient(c client.Client) {
	ts.client = c
}

func (ts *temporalService) Connect(ctx context.Context, params models.ConnectClientParams) error {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	if ts.connected {
		return nil
	}

	err := ts.client.Connect(ctx, params)
	if err != nil {
		return err
	}

	ts.connected = true
	return nil
}

func (ts *temporalService) ExecuteAsyncWorkflow(ctx context.Context, params models.ExecuteWorkflowParams) (models.WorkflowResponse, error) {
	return ts.client.ExecuteAsyncWorkflow(ctx, params)
}

func (ts *temporalService) ExecuteSyncWorkflow(ctx context.Context, params models.ExecuteWorkflowParams) (models.WorkflowResponse, error) {
	return ts.client.ExecuteSyncWorkflow(ctx, params)
}

func (ts *temporalService) Close(ctx context.Context) {
	ts.client.Close(ctx)
}

func (ts *temporalService) ListWorkflows(ctx context.Context, params models.ListWorkflowsParams) (models.ListWorkflowsResponse, error) {
	return ts.client.ListWorkflows(ctx, params)
}

func (ts *temporalService) GetWorkflowDetails(ctx context.Context, params models.GetWorkflowDetailsParams) (models.WorkflowDetailsResponse, error) {
	return ts.client.GetWorkflowDetails(ctx, params)
}

func (ts *temporalService) TerminateWorkflow(ctx context.Context, params models.TerminateWorkflowParams) error {
	return ts.client.TerminateWorkflow(ctx, params)
}

func (ts *temporalService) CancelWorkflow(ctx context.Context, params models.CancelWorkflowParams) error {
	return ts.client.CancelWorkflow(ctx, params)
}

func (ts *temporalService) SignalWorkflow(ctx context.Context, params models.SignalWorkflowParams) error {
	return ts.client.SignalWorkflow(ctx, params)
}

func (ts *temporalService) QueryWorkflow(ctx context.Context, params models.QueryWorkflowParams) (models.QueryWorkflowResponse, error) {
	return ts.client.QueryWorkflow(ctx, params)
}

func (ts *temporalService) GetNewWorker(ctx context.Context, params models.NewWorkerParams) (worker.Worker, error) {
	return worker.NewWorker(ctx, ts.client.(*client.TemporalClient), params)
}
