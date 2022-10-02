package model

type ProductRequest struct {
	Name      string  `json:"nome"`
	Price     float64 `json:"preco"`
	Available bool    `json:"disponivel"`
}

type ProductResponse struct {
	Name        string  `json:"nome"`
	Price       float64 `json:"preco"`
	Available   bool    `json:"disponivel"`
	Code        string  `json:"codigo_interno"`
	AddressShop string  `json:"endereco_loja"`
	BarCode     string  `json:"codigo_de_barras"`
}