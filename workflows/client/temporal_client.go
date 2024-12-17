package client

import (
	"context"

	"github.com/Zampfi/citadel/workflows/models"
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

func (tc *TemporalClient) StartAsyncWorkflow(ctx context.Context, params models.StartWorkflowParams) (models.StartWorkflowResponse, error) {
	wr, err := tc.baseTemporalClient.ExecuteWorkflow(ctx, params.ToTemporalStartWorkflowOptions(), params.Workflow, params.Args...)
	if err != nil {
		return models.StartWorkflowResponse{}, err
	}
	return models.StartWorkflowResponse{
		RunID:      wr.GetRunID(),
		WorkflowID: wr.GetID(),
		FirstRunID: wr.GetRunID(),
	}, nil
}

func (tc *TemporalClient) StartSyncWorkflow(ctx context.Context, params models.StartWorkflowParams) (models.StartWorkflowResponse, error) {
	wr, err := tc.baseTemporalClient.ExecuteWorkflow(ctx, params.ToTemporalStartWorkflowOptions(), params.Workflow, params.Args...)
	if err != nil {
		return models.StartWorkflowResponse{}, err
	}

	err = wr.Get(ctx, nil)
	if err != nil {
		return models.StartWorkflowResponse{}, err
	}

	return models.StartWorkflowResponse{
		RunID:      wr.GetRunID(),
		WorkflowID: wr.GetID(),
		FirstRunID: wr.GetRunID(),
	}, nil
}

func (tc *TemporalClient) Close(ctx context.Context) {
	tc.baseTemporalClient.Close()
}
