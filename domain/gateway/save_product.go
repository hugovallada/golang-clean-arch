package gateway

import (
	"github.com/hugovallada/go-clean-arch/domain"
)

type SaveProductGateway interface {
	Save(domain.Product) (string, error)
}
