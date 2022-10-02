package translator

import (
	"github.com/hugovallada/go-clean-arch/domain"
	"github.com/hugovallada/go-clean-arch/service/gateway/db/entity"
)

func FromProductToProductEntity(product domain.Product) entity.ProductEntity {
	return entity.ProductEntity{
		Name:        product.Name,
		Price:       product.Price,
		Available:   product.Available,
		Code:        product.Code,
		AddressShop: product.AddressShop,
		BarCode:     product.BarCode,
	}
}
