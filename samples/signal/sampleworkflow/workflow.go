package workflow

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

type SignalData struct {
	Message string
}

func SampleSignalWorkflow(ctx workflow.Context) error {
	logger := workflow.GetLogger(ctx)

	// Define the signal channel
	signalCh := workflow.GetSignalChannel(ctx, "sample-signal")

	// Define a variable to hold the signal data
	var signalData SignalData

	// Wait for the signal or context cancellation
	for {
		selector := workflow.NewSelector(ctx)

		// Add signal channel to the selector
		selector.AddReceive(signalCh, func(c workflow.ReceiveChannel, more bool) {
			c.Receive(ctx, &signalData)
			logger.Info("Received signal", "Message", signalData.Message)

			// Process the signal data
			if signalData.Message == "stop" {
				logger.Info("Stopping workflow as 'stop' signal received")
				return
			}
		})

		// Add context cancellation to the selector
		selector.AddFuture(workflow.NewTimer(ctx, time.Hour*24*365*100), func(f workflow.Future) {
			logger.Info("Workflow context canceled")
		})

		// Block until one of the selector's conditions is met
		selector.Select(ctx)

		// Simulate some work
		workflow.Sleep(ctx, 1*time.Second)
	}

}
