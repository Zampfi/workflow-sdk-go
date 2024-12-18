package models

import (
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
