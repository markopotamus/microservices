package grpc

import (
	"context"

	"github.com/markopotamus/microservices-proto/golang/order"
	"github.com/markopotamus/microservices/order/internal/application/core/domain"
)

func (a Adapter) Create(ctx context.Context, req *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	orderItems := []domain.OrderItem{}
	for _, oi := range req.OrderItems {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: oi.ProductCode,
			UnitPrice:   oi.UnitPrice,
			Quantity:    oi.Quantity,
		})
	}

	newOrder := domain.NewOrder(req.UserId, orderItems)
	result, err := a.api.PlaceOrder(newOrder)
	if err != nil {
		return nil, err
	}

	return &order.CreateOrderResponse{OrderId: result.ID}, nil
}
