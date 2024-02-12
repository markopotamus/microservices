package db

import (
	"fmt"

	"github.com/markopotamus/microservices/order/internal/application/core/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerID int64
	Status     string
	OrderItems []OrderItem
}

type OrderItem struct {
	gorm.Model
	ProductCode string
	UnitPrice   float32
	Quantity    int32
	OrderID     uint
}

type Adapter struct {
	db *gorm.DB
}

func NewAdapter(dataSrcUrl string) (*Adapter, error) {
	db, err := gorm.Open(mysql.Open(dataSrcUrl), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("db connection failed, err=%v", err)
	}

	err = db.AutoMigrate(&Order{}, &OrderItem{})
	if err != nil {
		return nil, fmt.Errorf("db migration failed, err=%v", err)
	}

	return &Adapter{db}, nil
}

func (a Adapter) Get(id string) (domain.Order, error) {
	var orderEntity Order

	res := a.db.First(&orderEntity, id)

	orderItems := []domain.OrderItem{}
	for _, oi := range orderEntity.OrderItems {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: oi.ProductCode,
			UnitPrice:   oi.UnitPrice,
			Quantity:    oi.Quantity,
		})
	}

	order := domain.Order{
		ID:         int64(orderEntity.ID),
		CustomerID: orderEntity.CustomerID,
		Status:     orderEntity.Status,
		OrderItems: orderItems,
		CreatedAt:  orderEntity.CreatedAt.UnixNano(),
	}

	return order, res.Error
}

func (a Adapter) Save(order *domain.Order) error {
	orderItemModels := []OrderItem{}
	for _, oi := range order.OrderItems {
		orderItemModels = append(orderItemModels, OrderItem{
			ProductCode: oi.ProductCode,
			UnitPrice:   oi.UnitPrice,
			Quantity:    oi.Quantity,
		})
	}

	orderModel := Order{
		CustomerID: order.CustomerID,
		Status:     order.Status,
		OrderItems: orderItemModels,
	}

	res := a.db.Create(&orderModel)
	if res.Error == nil {
		order.ID = int64(orderModel.ID)
	}

	return res.Error
}
