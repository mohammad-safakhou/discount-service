package core

import (
	"context"
	"discount-service/transport"
)

type DiscountFunctionalities interface {
	CreateDiscountRule(ctx context.Context, discountRule *DiscountRuleModel) error
	GetDiscountRule(ctx context.Context, accountId string) (DiscountRuleModel, error)
	ValidateDiscountCreationRules(ctx context.Context, input DiscountRuleModel) (bool, error)
	CalculateDiscount(ctx context.Context, accountId string, amount int) ()
}

type DiscountCoreContext struct {
	*transport.ApplicationContext
}