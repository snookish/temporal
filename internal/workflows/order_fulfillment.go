package workflows

import (
	"go.temporal.io/sdk/workflow"

	"github.com/snookish/temporal/internal/models"
)

func OrderFulfillmentWorkflow(ctx workflow.Context, cmd models.ProcessOrderCommand) (*models.OrderState, error) {
	return &models.OrderState{}, nil
}
