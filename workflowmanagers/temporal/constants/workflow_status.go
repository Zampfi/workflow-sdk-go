package constants

import "go.temporal.io/api/enums/v1"

type WorkflowExecutionStatus string

const (
	WorkflowExecutionStatusUnspecified    WorkflowExecutionStatus = "UNSPECIFIED"
	WorkflowExecutionStatusRunning        WorkflowExecutionStatus = "RUNNING"
	WorkflowExecutionStatusCompleted      WorkflowExecutionStatus = "COMPLETED"
	WorkflowExecutionStatusFailed         WorkflowExecutionStatus = "FAILED"
	WorkflowExecutionStatusCanceled       WorkflowExecutionStatus = "CANCELED"
	WorkflowExecutionStatusTerminated     WorkflowExecutionStatus = "TERMINATED"
	WorkflowExecutionStatusContinuedAsNew WorkflowExecutionStatus = "CONTINUED_AS_NEW"
	WorkflowExecutionStatusTimedOut       WorkflowExecutionStatus = "TIMED_OUT"
)

func (s *WorkflowExecutionStatus) FromTemporalWorkflowStatus(temporalStatus enums.WorkflowExecutionStatus) string {
	switch temporalStatus {
	case enums.WORKFLOW_EXECUTION_STATUS_RUNNING:
		return string(WorkflowExecutionStatusRunning)
	case enums.WORKFLOW_EXECUTION_STATUS_COMPLETED:
		return string(WorkflowExecutionStatusCompleted)
	case enums.WORKFLOW_EXECUTION_STATUS_FAILED:
		return string(WorkflowExecutionStatusFailed)
	case enums.WORKFLOW_EXECUTION_STATUS_CANCELED:
		return string(WorkflowExecutionStatusCanceled)
	case enums.WORKFLOW_EXECUTION_STATUS_TERMINATED:
		return string(WorkflowExecutionStatusTerminated)
	case enums.WORKFLOW_EXECUTION_STATUS_CONTINUED_AS_NEW:
		return string(WorkflowExecutionStatusContinuedAsNew)
	case enums.WORKFLOW_EXECUTION_STATUS_TIMED_OUT:
		return string(WorkflowExecutionStatusTimedOut)
	default:
		return string(WorkflowExecutionStatusUnspecified)
	}
}

func (s *WorkflowExecutionStatus) ToWorkflowExecutionStatus(status string) enums.WorkflowExecutionStatus {
	switch status {
	case string(WorkflowExecutionStatusRunning):
		return enums.WORKFLOW_EXECUTION_STATUS_RUNNING
	case string(WorkflowExecutionStatusCompleted):
		return enums.WORKFLOW_EXECUTION_STATUS_COMPLETED
	case string(WorkflowExecutionStatusFailed):
		return enums.WORKFLOW_EXECUTION_STATUS_FAILED
	case string(WorkflowExecutionStatusCanceled):
		return enums.WORKFLOW_EXECUTION_STATUS_CANCELED
	case string(WorkflowExecutionStatusTerminated):
		return enums.WORKFLOW_EXECUTION_STATUS_TERMINATED
	case string(WorkflowExecutionStatusContinuedAsNew):
		return enums.WORKFLOW_EXECUTION_STATUS_CONTINUED_AS_NEW
	case string(WorkflowExecutionStatusTimedOut):
		return enums.WORKFLOW_EXECUTION_STATUS_TIMED_OUT
	default:
		return enums.WORKFLOW_EXECUTION_STATUS_UNSPECIFIED
	}
}
