package main

import (
	"github.com/OzqurYalcin/google-analytics/src"
	"github.com/google/uuid"
)

func main() {
	api := new(ga.API)
	api.Lock()
	defer api.Unlock()
	client := new(ga.Client)
	client.ProtocolVersion = "1"
	client.TrackingID = "UA-xxxxxxxx-xx"
	client.DataSource = "web"
	client.ClientID = uuid.New().String()
	client.HitType = "transaction"

	product := new(ga.Product)
	product.SKU = "P1234"
	product.Name = "test"
	product.Brand = "test"
	product.Price = "25.00"
	product.Quantity = "1"
	product.Action = "purchase"
	client.Products = append(client.Products, product)

	client.TransactionID = "1111"
	client.TransactionRevenue = "25.00"
	client.TransactionTax = "0.00"

	client.CurrencyCode = "USD"
	api.Send(client)
}
