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

func GetAllProducts() ([]entity.ProductEntity, error) {
	conn, err := config.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.Query("SELECT * FROM products")

	if err != nil {
		return nil, err
	}

	var products []entity.ProductEntity

	for rows.Next() {
		var product entity.ProductEntity
		err = rows.Scan(&product.Id, &product.Name, &product.Price, &product.Available, &product.Code, &product.AddressShop, &product.BarCode)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func GetProductByCode(code string) (product entity.ProductEntity, err error) {
	conn, err := config.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sqlQuery := "Select * from products where code=$1"

	err = conn.QueryRow(sqlQuery, code).Scan(&product.Id, &product.Name, &product.Price, &product.Available, &product.Code, &product.AddressShop, &product.BarCode)

	if err != nil {
		return product, nil
	}

	return
}
