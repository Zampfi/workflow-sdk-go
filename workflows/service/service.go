package service

import (
	"context"
	"sync"

	"github.com/Zampfi/citadel/workflows/client"
	"github.com/Zampfi/citadel/workflows/models"
)

type WorkflowService interface {
	Connect(ctx context.Context, params models.ConnectClientParams) error
	StartAsyncWorkflow(ctx context.Context, params models.StartWorkflowParams) (models.StartWorkflowResponse, error)
	StartSyncWorkflow(ctx context.Context, params models.StartWorkflowParams) (models.StartWorkflowResponse, error)
	Close(ctx context.Context)
}

type workflowService struct {
	client    client.Client
	connected bool
	mu        sync.Mutex
}

func NewWorkflowService() WorkflowService {
	return &workflowService{
		client: &client.TemporalClient{},
	}
}

// This is to test mock client for testing only.
func (ws *workflowService) setClient(c client.Client) {
	ws.client = c
}

func (ws *workflowService) Connect(ctx context.Context, params models.ConnectClientParams) error {
	ws.mu.Lock()
	defer ws.mu.Unlock()

	if ws.connected {
		return nil
	}

	err := ws.client.Connect(ctx, params)
	if err != nil {
		return err
	}

	ws.connected = true
	return nil
}

func (ws *workflowService) StartAsyncWorkflow(ctx context.Context, params models.StartWorkflowParams) (models.StartWorkflowResponse, error) {
	return ws.client.StartAsyncWorkflow(ctx, params)
}

func (ws *workflowService) StartSyncWorkflow(ctx context.Context, params models.StartWorkflowParams) (models.StartWorkflowResponse, error) {
	return ws.client.StartSyncWorkflow(ctx, params)
}

func (ws *workflowService) Close(ctx context.Context) {
	ws.client.Close(ctx)
}
