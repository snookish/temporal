package workflows

import (
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"

	"github.com/snookish/temporal/internal/activities"
	"github.com/snookish/temporal/internal/models"
)

func OrderFulfillmentWorkflow(ctx workflow.Context, cmd models.ProcessOrderCommand) (*models.OrderState, error) {
	logger := workflow.GetLogger(ctx)
	logger.Info("Order fulfillment workflow started", "orderID", cmd.OrderID, "customerID", cmd.CustomerID)

	orderState := models.OrderState{
		OrderID: cmd.CustomerID,
		Status:  models.OrderStatusPending,
	}

	activityOpts := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 30,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second,
			BackoffCoefficient: 2.0, // Exponential backoff
			MaximumInterval:    time.Minute,
			MaximumAttempts:    5,
		},
	}

	ctx = workflow.WithActivityOptions(ctx, activityOpts)

	// STEP 1: Process Payment
	var processPaymentResp models.ProcessPaymentResponse
	if err := workflow.ExecuteActivity(ctx, activities.ProcessPaymentActivity, models.ProcessPaymentCommand{
		OrderID:       cmd.OrderID,
		Amount:        cmd.Amount,
		PaymentMethod: cmd.PaymentMethod,
	}).Get(ctx, &processPaymentResp); err != nil {
		orderState.Status = models.OrderStatusPaymentFailed
		return &orderState, err
	}

	orderState.Status = models.OrderStatusPaymentSucceeded
	orderState.CompensationSteps = append(orderState.CompensationSteps, models.CompensationStepRefund)
	logger.Info("Payment confirmed", "transactionID", processPaymentResp.TransactionID)

	return &orderState, nil
}
