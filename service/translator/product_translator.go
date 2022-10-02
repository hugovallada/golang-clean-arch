package translator

import (
	"github.com/google/uuid"
	"github.com/hugovallada/go-clean-arch/client/model"
	"github.com/hugovallada/go-clean-arch/domain"
)


func FromProductRequestToProduct(productRequest model.ProductRequest) domain.Product {
	return domain.Product{
		Name: productRequest.Name,
		Price: productRequest.Price,
		Available: productRequest.Available,
		Code: uuid.NewString(),
		AddressShop: "RP",
		BarCode: uuid.NewString(),
	}
}