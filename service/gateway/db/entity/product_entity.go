package entity

type ProductEntity struct {
	Id          int64
	Name        string
	Price       float64
	Available   bool
	Code        string
	AddressShop string
	BarCode     string
}
