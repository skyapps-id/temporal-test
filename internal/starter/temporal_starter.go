package starter

import (
	"context"
	"log"
	"time"

	"temporal-test/internal/workflow"

	"go.temporal.io/sdk/client"
)

type TemporalStarter struct {
	Client client.Client
}

func NewTemporalStarter() *TemporalStarter {
	c, err := client.Dial(client.Options{
		HostPort: "localhost:7233",
	})
	if err != nil {
		log.Fatal("Temporal dial error:", err)
	}
	return &TemporalStarter{Client: c}
}

func (t *TemporalStarter) StartPaymentWorkflow(orderID string, amount int64) error {
	workflowID := "payment_" + orderID

	input := workflow.PaymentInput{
		OrderID:     orderID,
		Amount:      amount,
		ReminderETA: time.Minute * 5,
		CancelETA:   time.Minute * 10,
	}

	_, err := t.Client.ExecuteWorkflow(
		context.Background(),
		client.StartWorkflowOptions{
			ID:        workflowID,
			TaskQueue: "PAYMENT_QUEUE",
		},
		workflow.PaymentWorkflow,
		input,
	)
	return err
}
