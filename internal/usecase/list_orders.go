package usecase

import (
	"context"
	"github.com/ramonsoterio/clean-architecture/internal/entity"
)

type ListOrdersOutputDTO []OrderOutputDTO
type ListOrdersUseCase struct {
	repository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(o entity.OrderRepositoryInterface) *ListOrdersUseCase {
	return &ListOrdersUseCase{repository: o}
}

func (l *ListOrdersUseCase) Execute(_ context.Context) (ListOrdersOutputDTO, error) {
	orders, err := l.repository.GetAll()
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
