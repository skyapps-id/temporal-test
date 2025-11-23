package workflow

import (
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

type PaymentInput struct {
	OrderID     string
	Amount      int64
	ReminderETA time.Duration // e.g. 10 minutes
	CancelETA   time.Duration // e.g. 20 minutes from start
}

func PaymentWorkflow(ctx workflow.Context, input PaymentInput) error {
	logger := workflow.GetLogger(ctx)
	logger.Info("Workflow started", "orderId", input.OrderID)

	// Activity options with retry policy
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 30,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts:    3,
			InitialInterval:    time.Second * 2,
			BackoffCoefficient: 2.0,
		},
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	// --- Step 1: Wait until reminder ETA ---
	logger.Info("Waiting until reminder ETA", "ETA", input.ReminderETA)
	if err := workflow.Sleep(ctx, input.ReminderETA); err != nil {
		return err
	}

	// --- Step 2: Send payment reminder ---
	if err := workflow.ExecuteActivity(ctx, SendPaymentReminderActivity, input.OrderID).Get(ctx, nil); err != nil {
		logger.Error("Failed to send reminder", "error", err)
	}

	// --- Step 3: Wait until cancel ETA (remaining duration) ---
	waitCancel := input.CancelETA - input.ReminderETA
	if waitCancel < 0 {
		waitCancel = 0
	}

	logger.Info("Waiting until cancel ETA", "ETA", waitCancel)
	if err := workflow.Sleep(ctx, waitCancel); err != nil {
		return err
	}

	// --- Step 4: Check payment status before cancelling ---
	var paid bool
	if err := workflow.ExecuteActivity(ctx, CheckPaymentStatusActivity, input.OrderID).Get(ctx, &paid); err != nil {
		logger.Error("Failed to check payment status", "error", err)
		return err
	}

	// Workflow ends successfully if paid
	if paid {
		logger.Info("Payment completed")
		return nil
	}

	// --- Step 5: Cancel order if not paid ---
	logger.Info("Payment not completed. Cancelling order")
	return workflow.ExecuteActivity(ctx, CancelOrderActivity, input.OrderID).Get(ctx, nil)
}
