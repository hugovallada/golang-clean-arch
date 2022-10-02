package translator

import (
	"github.com/google/uuid"
	"github.com/hugovallada/go-clean-arch/client/model"
	"github.com/hugovallada/go-clean-arch/domain"
	"github.com/hugovallada/go-clean-arch/service/gateway/db/entity"
)

func FromProductRequestToProduct(productRequest model.ProductRequest) domain.Product {
	return domain.Product{
		Name:        productRequest.Name,
		Price:       productRequest.Price,
		Available:   productRequest.Available,
		Code:        uuid.NewString(),
		AddressShop: "RP",
		BarCode:     uuid.NewString(),
	}
}

func FromEntityProductToProduct(product entity.ProductEntity) domain.Product {
	return domain.Product{
		Name:        product.Name,
		Price:       product.Price,
		Available:   product.Available,
		Code:        product.Code,
		AddressShop: product.AddressShop,
		BarCode:     product.BarCode,
	}
}

func FromListOfEntityProductsToListOfProducts(productsEntity []entity.ProductEntity) (products []domain.Product) {
	for _, product := range productsEntity {
		products = append(products, FromEntityProductToProduct(product))
	}
	return products
}

func FromProductToProductResponse(product domain.Product) model.ProductResponse {
	return model.ProductResponse{
		Name:        product.Name,
		Price:       product.Price,
		Available:   product.Available,
		Code:        product.Code,
		AddressShop: product.AddressShop,
		BarCode:     product.BarCode,
	}
}

func FromListOfProductsToListOfProductsReponse(products []domain.Product) (productsResponse []model.ProductResponse) {
	for _, product := range products {
		productsResponse = append(productsResponse, FromProductToProductResponse(product))
	}
	return
}
