package client

import (
	"context"

	"github.com/Zampfi/citadel/workflows/temporal/models"
)

type Client interface {
	Connect(ctx context.Context, params models.ConnectClientParams) error
	ExecuteAsyncWorkflow(ctx context.Context, params models.ExecuteWorkflowParams) (models.WorkflowResponse, error)
	ExecuteSyncWorkflow(ctx context.Context, params models.ExecuteWorkflowParams) (models.WorkflowResponse, error)
	Close(ctx context.Context)
}
