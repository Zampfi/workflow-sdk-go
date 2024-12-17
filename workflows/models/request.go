package models

import (
	"go.temporal.io/sdk/client"
)

type ConnectClientParams struct {
	Options ConnectClientOptions
}

type ConnectClientOptions = client.Options

type StartWorkflowParams struct {
	Options  StartWorkflowOptions
	Workflow interface{}
	Args     []interface{}
}

type StartWorkflowOptions = client.StartWorkflowOptions
