package models

import (
	v1 "go.temporal.io/api/common/v1"
	historypb "go.temporal.io/api/history/v1"
)

type WorkflowResponse struct {
	RunID      string
	WorkflowID string
}

type WorkflowDetailsResponse struct {
	Details      WorkflowExecutionDetails
	EventHistory []*Event
}

type WorkflowExecutionDetails struct {
	WorkflowID       string
	RunID            string
	WorkflowType     string
	TaskQueue        string
	Status           string
	StartTime        int64
	CloseTime        int64
	ExecutionTime    int64
	HistoryLength    int64
	SearchAttributes *v1.SearchAttributes
}

type Event struct {
	historypb.HistoryEvent
}

type ListWorkflowsResponse struct {
	Workflows     []WorkflowExecutionDetails
	NextPageToken []byte
}
