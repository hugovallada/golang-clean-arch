package usecase

import (
	"github.com/hugovallada/go-clean-arch/domain"
	"github.com/hugovallada/go-clean-arch/domain/gateway"
)

func SaveProductUseCase(product domain.Product, save_product_gateway gateway.SaveProductGateway) (string, error) {
	return save_product_gateway.Save(product)
}
