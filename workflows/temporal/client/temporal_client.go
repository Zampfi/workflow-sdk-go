package client

import (
	"context"

	"github.com/Zampfi/citadel/workflows/temporal/models"
	"go.temporal.io/sdk/client"
)

type TemporalClient struct {
	baseTemporalClient client.Client
}

func (tc *TemporalClient) Connect(ctx context.Context, params models.ConnectClientParams) error {
	newClient, err := client.Dial(params.ToTemporalClientOptions())
	if err != nil {
		return err
	}
	tc.baseTemporalClient = newClient
	return nil
}

func (tc *TemporalClient) ExecuteAsyncWorkflow(ctx context.Context, params models.ExecuteWorkflowParams) (models.WorkflowResponse, error) {
	wr, err := tc.baseTemporalClient.ExecuteWorkflow(ctx, params.ToTemporalStartWorkflowOptions(), params.Workflow, params.Args...)
	if err != nil {
		return models.WorkflowResponse{}, err
	}
	return models.WorkflowResponse{
		RunID:      wr.GetRunID(),
		WorkflowID: wr.GetID(),
		FirstRunID: wr.GetRunID(),
	}, nil
}

func (tc *TemporalClient) ExecuteSyncWorkflow(ctx context.Context, params models.ExecuteWorkflowParams) (models.WorkflowResponse, error) {
	wr, err := tc.baseTemporalClient.ExecuteWorkflow(ctx, params.ToTemporalStartWorkflowOptions(), params.Workflow, params.Args...)
	if err != nil {
		return models.WorkflowResponse{}, err
	}

	err = wr.Get(ctx, params.ResultPtr)
	if err != nil {
		return models.WorkflowResponse{}, err
	}

	return models.WorkflowResponse{
		RunID:      wr.GetRunID(),
		WorkflowID: wr.GetID(),
		FirstRunID: wr.GetRunID(),
	}, nil
}

func (tc *TemporalClient) Close(ctx context.Context) {
	tc.baseTemporalClient.Close()
}
