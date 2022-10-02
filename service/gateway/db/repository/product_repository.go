package repository

import (
	"github.com/hugovallada/go-clean-arch/service/gateway/db/config"
	"github.com/hugovallada/go-clean-arch/service/gateway/db/entity"
)

func SaveProduct(product entity.ProductEntity) (string, error) {
	conn, err := config.OpenConnection()
	if err != nil {
		return "", err
	}
	defer conn.Close()

	var code string

	sql := "INSERT INTO products (name, price, availability, code, address, bar_code) VALUES ($1, $2, $3, $4, $5, $6) RETURNING code"

	err = conn.QueryRow(sql, product.Name, product.Price, product.Available, product.Code, product.AddressShop, product.BarCode).Scan(&code)

	return code, err
}
