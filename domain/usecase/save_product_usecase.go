package usecase

import (
	"github.com/hugovallada/go-clean-arch/domain"
	errormodel "github.com/hugovallada/go-clean-arch/domain/error_model"
	"github.com/hugovallada/go-clean-arch/domain/gateway"
)

func SaveProductUseCase(product domain.Product,
	saveProductGateway gateway.SaveProductGateway,
	getProductGateway gateway.GetProductsGateway,
) (string, error) {
	savedProduct, err := getProductGateway.GetByCode(product.Code)
	if err != nil {
		return "", err
	}

	if !isEmpty(savedProduct) {
		return "", &errormodel.UnprocessableEntityError{
			Message: "Esse produto já existe",
		}
	}

	return saveProductGateway.Save(product)
}

func isEmpty[T comparable](object T) bool {
	var emptyStruct T
	return object == emptyStruct
}
