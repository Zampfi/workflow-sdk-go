package models

// Output structs for each method

type StartWorkflowResponse struct {
	RunID      string
	WorkflowID string
	FirstRunID string
}
