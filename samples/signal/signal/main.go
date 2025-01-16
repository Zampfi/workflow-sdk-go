package main

import (
	"context"
	"fmt"

	signalsampleworkflow "github.com/Zampfi/workflow-sdk-go/samples/signal/sampleworkflow"
	"github.com/Zampfi/workflow-sdk-go/workflowmanagers/temporal"
	"github.com/Zampfi/workflow-sdk-go/workflowmanagers/temporal/models"
)

func main() {
	ctx := context.Background()
	temporalService := temporal.NewTemporalService()
	err := temporalService.Connect(ctx, models.ConnectClientParams{})
	if err != nil {
		panic(err)
	}

	err = temporalService.SignalWorkflow(ctx, models.SignalWorkflowParams{
		WorkflowID: "SampleSignalWorkflow-NPh4SjCj28Nm3rNYgDNFn4", //pass the workflowID of the workflow you want to signal
		SignalName: "sample-signal",
		Arg:        signalsampleworkflow.SignalData{Message: "stop"},
	})
	if err != nil {
		fmt.Printf("Error sending signal: %v\n", err)
	}
}
