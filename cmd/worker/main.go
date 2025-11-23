package main

import (
	"log"

	wf "temporal-test/internal/workflow"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.Dial(client.Options{
		HostPort: "localhost:7233",
	})
	if err != nil {
		log.Fatalln("Unable to connect to Temporal:", err)
	}
	defer c.Close()

	w := worker.New(c, "PAYMENT_QUEUE", worker.Options{})

	w.RegisterWorkflow(wf.PaymentWorkflow)
	w.RegisterActivity(wf.SendPaymentReminderActivity)
	w.RegisterActivity(wf.CheckPaymentStatusActivity)
	w.RegisterActivity(wf.CancelOrderActivity)

	log.Println("Worker started...")
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatal("unable to start worker", err)
	}
}
