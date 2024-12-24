package workflow

import (
	temporalworkflow "go.temporal.io/sdk/workflow"
)

type Workflow struct {
	Function        interface{}
	RegisterOptions RegisterWorkflowOptions
}

type RegisterWorkflowOptions = temporalworkflow.RegisterOptions
