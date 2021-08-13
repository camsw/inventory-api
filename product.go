package main

import (
	"encoding/json"
	"io"
)

// The structure for our API product
type Product struct {
	ID            int64   `json:"id"`
	Description   string  `json:"description"`
	PurchaseDate  string  `json:"purchase_date"`
	SalesDate     string  `json:"sales_date"`
	ShipDate      string  `json:"ship_date"`
	PurchasePrice float64 `json:"purchase_price"`
	SalesPrice    float64 `json:"sales_price"`
}

// A collection of Product objects
type Products []*Product

func (p *Product) fromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func (p *Products) toJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func getProducts() Products {
	return productList
}

var productList = []*Product{
	&Product{
		ID:            1,
		Description:   "7 Diamonds - Closer To You Fannel - Blue - size M",
		PurchaseDate:  "2020-10-03",
		SalesDate:     "2021-03-25",
		ShipDate:      "2021-03-26",
		PurchasePrice: 12.75,
		SalesPrice:    49,
	},
	&Product{
		ID:            2,
		Description:   "7 For All Mankind - Slimmy Slim - Everglade - size 34",
		PurchaseDate:  "2021-01-21",
		SalesDate:     "2021-07-30",
		ShipDate:      "2021-07-31",
		PurchasePrice: 22.75,
		SalesPrice:    56.90,
	},
}
