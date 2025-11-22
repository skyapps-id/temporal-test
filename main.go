package main

import (
	"context"
	"log"
	"time"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

const (
	QueueName    = "PAYMENT_QUEUE"
	WorkflowName = "PaymentETAWorkflow"
)

func main() {
	// Mode pilih: worker atau starter
	runAsWorker := true

	if runAsWorker {
		startWorker()
	} else {
		startWorkflow()
	}
}

func startWorker() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatal("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, QueueName, worker.Options{})

	w.RegisterWorkflow(PaymentETAWorkflow)
	w.RegisterActivity(SendAlertActivity)

	log.Println("Worker started...")
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatal("Unable to start worker", err)
	}
}

func startWorkflow() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatal("Unable to create client", err)
	}
	defer c.Close()

	input := PaymentWorkflowInput{
		PaymentID: "INV-001",
		ETA:       10 * time.Second,
	}

	we, err := c.ExecuteWorkflow(context.Background(), client.StartWorkflowOptions{
		ID:        "payment-eta-workflow-1",
		TaskQueue: QueueName,
	}, PaymentETAWorkflow, input)

	if err != nil {
		log.Fatal("Unable to start workflow", err)
	}

	log.Println("Started workflow:", "WorkflowID", we.GetID(), "RunID", we.GetRunID())
}
