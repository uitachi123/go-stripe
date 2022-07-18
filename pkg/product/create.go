package product

import (
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/product"
	"go.uber.org/zap"
)

func Create(logger *zap.Logger) (*stripe.Product, error) {
	// Structured context as strongly typed Field values.
	starter_subscription := &stripe.ProductParams{
		Name:        stripe.String("Starter Subscription"),
		Description: stripe.String("$12/Month subscription"),
	}

	starter_product, err := product.New(starter_subscription)
	if err != nil {
		return nil, err
	}

	logger.Debug("Success", zap.String("product ID", starter_product.ID))
	return starter_product, nil
}
