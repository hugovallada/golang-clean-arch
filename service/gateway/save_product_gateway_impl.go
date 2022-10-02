package gateway

import (
	"github.com/hugovallada/go-clean-arch/domain"
	"github.com/hugovallada/go-clean-arch/service/gateway/db/repository"
	"github.com/hugovallada/go-clean-arch/service/gateway/db/translator"
)

type SaveProductGatewayImpl struct {
}

func (s *SaveProductGatewayImpl) Save(product domain.Product) (string, error) {
	entity := translator.FromProductToProductEntity(product)
	return repository.SaveProduct(entity)
}
