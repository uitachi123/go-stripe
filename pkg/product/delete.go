package product

import (
	"github.com/stripe/stripe-go/v72/product"
	"go.uber.org/zap"
)

func Delete(id string, logger *zap.Logger) error {
	p, err := product.Del(id, nil)
	if err != nil {
		return err
	}
	logger.Debug("Success", zap.String("Deleted product:", p.ID))
	return nil
}
