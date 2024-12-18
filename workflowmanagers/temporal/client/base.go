package client

import (
	"context"

	"github.com/Zampfi/citadel/workflowmanagers/temporal/models"
)

type Client interface {
	Connect(ctx context.Context, params models.ConnectClientParams) error
	ExecuteAsyncWorkflow(ctx context.Context, params models.ExecuteWorkflowParams) (models.WorkflowResponse, error)
	ExecuteSyncWorkflow(ctx context.Context, params models.ExecuteWorkflowParams) (models.WorkflowResponse, error)
	ListWorkflows(ctx context.Context, params models.ListWorkflowsParams) (models.ListWorkflowsResponse, error)
	GetWorkflowDetails(ctx context.Context, params models.GetWorkflowDetailsParams) (models.WorkflowDetailsResponse, error)
	TerminateWorkflow(ctx context.Context, params models.TerminateWorkflowParams) error
	CancelWorkflow(ctx context.Context, params models.CancelWorkflowParams) error
	Close(ctx context.Context)
}
