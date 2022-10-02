package usecase

import (
	"github.com/hugovallada/go-clean-arch/domain"
	"github.com/hugovallada/go-clean-arch/domain/gateway"
)

func GetAllProducts(getProductGateway gateway.GetProductsGateway) ([]domain.Product, error) {
	return getProductGateway.GetAll()
}
