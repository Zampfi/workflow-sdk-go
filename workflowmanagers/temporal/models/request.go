package models

import (
	activity "github.com/Zampfi/citadel/workflowmanagers/temporal/activity"
	workflow "github.com/Zampfi/citadel/workflowmanagers/temporal/workflow"
	commonpb "go.temporal.io/api/common/v1"
	enumspb "go.temporal.io/api/enums/v1"
	"go.temporal.io/sdk/client"
)

type ConnectClientParams struct {
	Options ConnectClientOptions
}

type ConnectClientOptions = client.Options

type ExecuteWorkflowParams struct {
	Options   StartWorkflowOptions
	Workflow  interface{}
	Args      []interface{}
	ResultPtr interface{}
}

type StartWorkflowOptions = client.StartWorkflowOptions

type GetWorkflowDetailsParams struct {
	WorkflowID string
	RunID      string
}

type ListWorkflowsParams struct {
	PageSize      int32
	Namespace     string
	NextPageToken []byte
	Query         string
}

type CancelWorkflowParams struct {
	WorkflowID string
	RunID      string
}

type TerminateWorkflowParams struct {
	WorkflowID string
	RunID      string
	Reason     string
	Details    []interface{}
}

type SignalWorkflowParams struct {
	WorkflowID string
	RunID      string
	SignalName string
	Arg        interface{}
}

type QueryWorkflowParams struct {
	WorkflowID           string
	RunID                string
	QueryType            string
	Args                 []interface{}
	QueryRejectCondition enumspb.QueryRejectCondition
	Header               *commonpb.Header
}

type NewWorkerParams struct {
	TaskQueue     string
	Workflows     []workflow.Workflow
	Activities    []activity.Activity
	Options       WorkerOptions
	RegisterTasks bool
}

type RunWorkerParams struct {
}
