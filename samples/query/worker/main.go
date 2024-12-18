package main

import (
	"log"

	querysampleworkflow "github.com/Zampfi/citadel/samples/query/sampleworkflow"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	// Create a Temporal client
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	// Create a worker
	w := worker.New(c, "test_query", worker.Options{})

	// Register the workflow
	w.RegisterWorkflow(querysampleworkflow.SampleWorkflow)

	// Start the worker
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
