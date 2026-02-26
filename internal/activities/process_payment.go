package activities

import (
	"context"
	"fmt"

	"go.temporal.io/sdk/activity"

	"github.com/snookish/temporal/internal/models"
)

func ProcessPaymentActivity(ctx context.Context, cmd models.ProcessPaymentCommand) (*models.ProcessPaymentResponse, error) {
	logger := activity.GetLogger(ctx)
	activityInfo := activity.GetInfo(ctx)

	idempotencyKey := fmt.Sprintf("%s-%s", activityInfo.WorkflowExecution.ID, activityInfo.ActivityID)

	logger.Info(
		"Payment processing activity started",
		"idempotencyKey", idempotencyKey,
		"orderID", cmd.OrderID, "paymentMethod",
		cmd.PaymentMethod, "activityID", activityInfo.ActivityID,
		"amount", cmd.Amount.Amount, "currency", cmd.Amount.Currency,
	)

	// Check if we already processed this payment
	// existingPayment := a.paymentDB.GetByIdempotencyKey(idempotencyKey)
	// if existingPayment != nil {
	//     logger.Info("Payment already processed, returning cached result")
	//     return existingPayment, nil
	// }

	// Process payment with payment provider
	// result := a.paymentProvider.Charge(paymentMethod, amount, idempotencyKey)

	response := models.ProcessPaymentResponse{
		TransactionID: idempotencyKey,
		Status:        models.TransactionStatusSuccess,
	}

	return &response, nil
}
