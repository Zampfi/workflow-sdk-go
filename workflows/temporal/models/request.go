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
