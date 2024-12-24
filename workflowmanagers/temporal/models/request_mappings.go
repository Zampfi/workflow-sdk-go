package models

import (
	"github.com/jinzhu/copier"
	"go.temporal.io/api/enums/v1"
	"go.temporal.io/api/workflowservice/v1"
	"go.temporal.io/sdk/client"
	temporalworker "go.temporal.io/sdk/worker"
)

func (c *ConnectClientParams) ToTemporalClientOptions() client.Options {
	options := client.Options{}
	copier.Copy(&options, &c.Options)
	return options
}

func (s *ExecuteWorkflowParams) ToTemporalStartWorkflowOptions() client.StartWorkflowOptions {
	options := client.StartWorkflowOptions{}
	copier.Copy(&options, s.Options)
	options.WorkflowIDReusePolicy = enums.WORKFLOW_ID_REUSE_POLICY_REJECT_DUPLICATE
	options.WorkflowIDConflictPolicy = enums.WORKFLOW_ID_CONFLICT_POLICY_FAIL
	return options
}

func (s *ListWorkflowsParams) ToTemporalListWorkflowsRequest() *workflowservice.ListWorkflowExecutionsRequest {
	request := &workflowservice.ListWorkflowExecutionsRequest{}
	copier.Copy(request, s)
	return request
}

func (s *QueryWorkflowParams) ToTemporalQueryWorkflowWithOptionsRequest() *client.QueryWorkflowWithOptionsRequest {
	request := &client.QueryWorkflowWithOptionsRequest{}
	copier.Copy(&request, s)
	return request
}

func (s *NewWorkerParams) ToTemporalWorkerOptions() temporalworker.Options {
	options := temporalworker.Options{}
	copier.Copy(&options, s.Options)
	return options
}
