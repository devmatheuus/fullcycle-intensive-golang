package usecase

import "github.com/devmatheuus/pfa-go/internal/order/entity"

type OrderInputDTO struct {
	ID    string
	Price float64
	Tax   float64
}

type OrderOutputDTO struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

type CalculateFinalPriceUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewCalculateFinalPriceUseCase(orderRepository entity.OrderRepositoryInterface) *CalculateFinalPriceUseCase {
	return &CalculateFinalPriceUseCase{
		OrderRepository: orderRepository,
	}
}

func (c *CalculateFinalPriceUseCase) Exec(order *OrderInputDTO) (*OrderOutputDTO, error) {
	newOrder, err := entity.NewOrder(order.ID, order.Price, order.Tax)
	if err != nil {
		return nil, err
	}

	err = newOrder.CalculateFinalPrice()
	if err != nil {
		return nil, err
	}

	err = c.OrderRepository.Save(newOrder)
	if err != nil {
		return nil, err
	}

	return &OrderOutputDTO{
		ID:         newOrder.ID,
		Price:      newOrder.Price,
		Tax:        newOrder.Tax,
		FinalPrice: newOrder.FinalPrice,
	}, nil
}
