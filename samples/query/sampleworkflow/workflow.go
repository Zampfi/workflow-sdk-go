package sampleworkflow

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func SampleWorkflow(ctx workflow.Context) error {
	var counter int

	err := workflow.SetQueryHandler(ctx, "getCounter", func() (int, error) {
		return counter, nil
	})
	if err != nil {
		return err
	}

	for i := 0; i < 10; i++ {
		counter++
		workflow.Sleep(ctx, time.Second)
	}

	return nil
}
