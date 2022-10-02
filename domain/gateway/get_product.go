package gateway

import "github.com/hugovallada/go-clean-arch/domain"

type GetProductsGateway interface {
	GetAll() ([]domain.Product, error)
	GetByCode(code string) (domain.Product, error)
}
