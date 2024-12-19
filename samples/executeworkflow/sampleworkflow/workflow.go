package sampleworkflow

import (
	"context"
	"time"

	"go.temporal.io/sdk/workflow"
)

func SampleWorkflow(ctx workflow.Context, number int) (int, error) {

	// This is a sample workflow that executes a single activity
	// and returns the result of the activity
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
		TaskQueue:           "test_execute",
	}

	ctx = workflow.WithActivityOptions(ctx, ao)

	var result int
	err := workflow.ExecuteActivity(ctx, SampleActivity, number).Get(ctx, &result)
	if err != nil {
		return 0, err
	}

	return result, nil
}

func SampleActivity(ctx context.Context, number int) (int, error) {
	return number * 100, nil
}
