package sampleworkflow

import (
	"go.temporal.io/sdk/workflow"
)

func SampleWorkflow(ctx workflow.Context, number int) (int, error) {
	return number * 2, nil
}
