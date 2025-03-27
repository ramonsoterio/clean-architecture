package usecase

import (
	"context"
	"github.com/ramonsoterio/clean-architecture/internal/entity"
)

type ListOrdersOutputDTO []OrderOutputDTO
type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(OrderRepository entity.OrderRepositoryInterface) *ListOrdersUseCase {
	return &ListOrdersUseCase{OrderRepository: OrderRepository}
}

func (l *ListOrdersUseCase) Execute(_ context.Context) (ListOrdersOutputDTO, error) {
	orders, err := l.OrderRepository.GetAll()
	if err != nil {
		return ListOrdersOutputDTO{}, err
	}
	outputOrders := mapOrdersToDTOs(orders)
	return outputOrders, nil
}

func mapOrdersToDTOs(orders []entity.Order) (dtos ListOrdersOutputDTO) {
	for _, order := range orders {
		dtos = append(dtos, OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}
	return
}
