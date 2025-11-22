package main

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

type PaymentWorkflowInput struct {
	PaymentID string
	ETA       time.Duration // contoh: 30 * time.Minute
}

func PaymentETAWorkflow(ctx workflow.Context, input PaymentWorkflowInput) error {
	logger := workflow.GetLogger(ctx)
	logger.Info("Workflow started", "payment_id", input.PaymentID)

	// Set activity options
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute * 1,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	// Sleep sesuai ETA — Temporal yang urus (bukan code)
	logger.Info("Sleeping until ETA is reached...")
	err := workflow.Sleep(ctx, input.ETA)
	if err != nil {
		return err
	}

	// After ETA → Kirim alert activity
	logger.Info("ETA reached, sending alert...")
	return workflow.ExecuteActivity(ctx, SendAlertActivity, input.PaymentID).Get(ctx, nil)
}
