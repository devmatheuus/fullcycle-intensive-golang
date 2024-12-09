package entity_test

import (
	"testing"

	"github.com/devmatheuus/pfa-go/internal/order/entity"
	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyIdWhenCreateANewOrderThenShouldReceiveAnError(t *testing.T) {
	order := entity.Order{}
	assert.Error(t, order.IsValid(), "invalid id")
}

func TestGivenAnEmptyPriceWhenCreateANewOrderThenShouldReceiveAnError(t *testing.T) {
	order := entity.Order{ID: "1"}
	assert.Error(t, order.IsValid(), "invalid price")
}

func TestGivenAValidParamsWhenCreateANewOrderThenShouldReceiveCreateOrderWithAllParams(t *testing.T) {
	order, _ := entity.NewOrder("1", 10.0, 1.0)
	assert.NotNil(t, order)
	assert.Equal(t, "1", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 1.0, order.Tax)
}

func TestGivenAValidParamsWhenCallCalculateFinalPriceThenShouldReturnFinalPrice(t *testing.T) {
	order, _ := entity.NewOrder("1", 10, 2)
	order.CalculateFinalPrice()
	assert.Equal(t, 12.0, order.FinalPrice)
}
