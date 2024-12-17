package client

import (
	"context"

	"github.com/Zampfi/citadel/workflows/models"
)

type Client interface {
	Connect(ctx context.Context, params models.ConnectClientParams) error
	StartAsyncWorkflow(ctx context.Context, params models.StartWorkflowParams) (models.StartWorkflowResponse, error)
	StartSyncWorkflow(ctx context.Context, params models.StartWorkflowParams) (models.StartWorkflowResponse, error)
	Close(ctx context.Context)
}
