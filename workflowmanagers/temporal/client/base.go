package client

import (
	"context"

	"github.com/Zampfi/workflow-sdk-go/workflowmanagers/temporal/models"
)

type Client interface {
	Connect(ctx context.Context, params models.ConnectClientParams) error
	ExecuteAsyncWorkflow(ctx context.Context, params models.ExecuteWorkflowParams) (models.WorkflowResponse, error)
	ExecuteSyncWorkflow(ctx context.Context, params models.ExecuteWorkflowParams) (models.WorkflowResponse, error)
	ExecuteScheduledWorkflow(ctx context.Context, params models.ExecuteWorkflowWithScheduleParams) (models.ScheduledWorkflowResponse, error)
	ListWorkflows(ctx context.Context, params models.ListWorkflowsParams) (models.ListWorkflowsResponse, error)
	GetWorkflowDetails(ctx context.Context, params models.GetWorkflowDetailsParams) (models.WorkflowDetailsResponse, error)
	TerminateWorkflow(ctx context.Context, params models.TerminateWorkflowParams) error
	CancelWorkflow(ctx context.Context, params models.CancelWorkflowParams) error
	SignalWorkflow(ctx context.Context, params models.SignalWorkflowParams) error
	QueryWorkflow(ctx context.Context, params models.QueryWorkflowParams) (models.QueryWorkflowResponse, error)
	QuerySchedule(ctx context.Context, params models.QueryScheduleParams) (models.QueryScheduleResponse, error)
	Close(ctx context.Context)
}
