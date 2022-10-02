package gateway

import (
	"github.com/hugovallada/go-clean-arch/domain"
	"github.com/hugovallada/go-clean-arch/service/gateway/db/repository"
	"github.com/hugovallada/go-clean-arch/service/translator"
)

type GetProductsGatewayImpl struct{}

func (gp *GetProductsGatewayImpl) GetAll() ([]domain.Product, error) {
	products, err := repository.GetAllProducts()
	if err != nil {
		return nil, err
	}

	return translator.FromListOfEntityProductsToListOfProducts(products), nil
}

func (gp *GetProductsGatewayImpl) GetByCode(code string) (product domain.Product, err error) {
	productEntity, err := repository.GetProductByCode(code)

	if err != nil {
		return
	}

	return translator.FromEntityProductToProduct(productEntity), nil
}
