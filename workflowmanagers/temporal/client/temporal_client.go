package client

import (
	"context"

	"github.com/Zampfi/citadel/workflowmanagers/temporal/models"
	"go.temporal.io/api/enums/v1"
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
	}, nil
}

func (tc *TemporalClient) Close(ctx context.Context) {
	tc.baseTemporalClient.Close()
}

func (tc *TemporalClient) GetWorkflowDetails(ctx context.Context, params models.GetWorkflowDetailsParams) (models.WorkflowDetailsResponse, error) {
	wr := tc.baseTemporalClient.GetWorkflow(ctx, params.WorkflowID, params.RunID)
	if wr == nil {
		return models.WorkflowDetailsResponse{}, nil
	}

	workflowExecutionDetails, err := tc.getWorkflowExecutionDetails(ctx, params.WorkflowID, wr.GetRunID())
	if err != nil {
		return models.WorkflowDetailsResponse{}, err
	}

	workflowExecutionHistory, err := tc.getWorkflowExecutionHistory(ctx, params.WorkflowID, wr.GetRunID())
	if err != nil {
		return models.WorkflowDetailsResponse{}, err
	}

	return models.WorkflowDetailsResponse{
		Details:      workflowExecutionDetails,
		EventHistory: workflowExecutionHistory,
	}, nil
}

func (tc *TemporalClient) getWorkflowExecutionDetails(ctx context.Context, workflowID, runID string) (models.WorkflowExecutionDetails, error) {
	workflowExecutionInfo, err := tc.baseTemporalClient.DescribeWorkflowExecution(ctx, workflowID, runID)
	if err != nil {
		return models.WorkflowExecutionDetails{}, err
	}

	if workflowExecutionInfo == nil {
		return models.WorkflowExecutionDetails{}, nil
	}

	workflowExecutionDetails := models.WorkflowExecutionDetails{}
	workflowExecutionDetails.FromDescribeWorkflowResponse(workflowExecutionInfo.WorkflowExecutionInfo)
	return workflowExecutionDetails, nil
}

func (tc *TemporalClient) getWorkflowExecutionHistory(ctx context.Context, workflowID, runID string) ([]*models.Event, error) {
	historyIterator := tc.baseTemporalClient.GetWorkflowHistory(ctx, workflowID, runID, false, enums.HISTORY_EVENT_FILTER_TYPE_ALL_EVENT)
	events := []*models.Event{}

	for historyIterator.HasNext() {
		currentEvent, err := historyIterator.Next()
		if err != nil {
			return nil, err
		}

		event := &models.Event{}
		event.FromHistoryEvent(currentEvent)
		events = append(events, event)
	}

	return events, nil
}

func (tc *TemporalClient) ListWorkflows(ctx context.Context, params models.ListWorkflowsParams) (models.ListWorkflowsResponse, error) {
	workflowsList, err := tc.baseTemporalClient.ListWorkflow(ctx, params.ToTemporalListWorkflowsRequest())
	if err != nil {
		return models.ListWorkflowsResponse{}, err
	}

	if workflowsList == nil {
		return models.ListWorkflowsResponse{}, nil
	}

	workflows := []models.WorkflowExecutionDetails{}
	for _, workflowInfo := range workflowsList.Executions {
		workflowExecutionDetails := models.WorkflowExecutionDetails{}
		workflowExecutionDetails.FromDescribeWorkflowResponse(workflowInfo)
		workflows = append(workflows, workflowExecutionDetails)
	}

	return models.ListWorkflowsResponse{
		Workflows:     workflows,
		NextPageToken: workflowsList.NextPageToken,
	}, nil
}

func (tc *TemporalClient) TerminateWorkflow(ctx context.Context, params models.TerminateWorkflowParams) error {
	return tc.baseTemporalClient.TerminateWorkflow(ctx, params.WorkflowID, params.RunID, params.Reason, params.Details...)
}

func (tc *TemporalClient) CancelWorkflow(ctx context.Context, params models.CancelWorkflowParams) error {
	return tc.baseTemporalClient.CancelWorkflow(ctx, params.WorkflowID, params.RunID)
}

func (tc *TemporalClient) SignalWorkflow(ctx context.Context, params models.SignalWorkflowParams) error {
	return tc.baseTemporalClient.SignalWorkflow(ctx, params.WorkflowID, params.RunID, params.SignalName, params.Arg)
}

func (tc *TemporalClient) QueryWorkflow(ctx context.Context, params models.QueryWorkflowParams) (models.QueryWorkflowResponse, error) {
	queryResponse, err := tc.baseTemporalClient.QueryWorkflowWithOptions(ctx, params.ToTemporalQueryWorkflowWithOptionsRequest())
	if err != nil {
		return models.QueryWorkflowResponse{}, err
	}

	response := models.QueryWorkflowResponse{}
	response.FromQueryWorkflowWithOptionsResponse(queryResponse)
	return response, nil
}
