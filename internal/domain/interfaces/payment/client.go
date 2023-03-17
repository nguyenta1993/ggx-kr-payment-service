package payment

import (
	"context"
	"github.com/gogovan/ggx-kr-payment-service/internal/domain/entity"
	"net/http"
)

type PaymentClient interface {
	Checkout(ctx context.Context, transaction entity.Transaction) (string, http.Header, error)
}
