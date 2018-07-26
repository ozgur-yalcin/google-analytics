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
	client.HitType = "pageview"
	client.DocumentHostName = "example.com"
	client.DocumentPath = "/payment"
	client.DocumentTitle = "Payment"
	client.ClientID = uuid.New().String()

	product := new(ga.Product)
	product.Action = "purchase"
	product.SKU = "P1234"
	product.Name = "name"
	product.Brand = "brand"
	product.Category = "category"
	product.Variant = "variant"
	product.Price = "25.00"
	product.Quantity = "1"
	product.Position = "1"
	client.Products = append(client.Products, product)

	client.TransactionID = "T1234"
	client.TransactionAffiliation = "affiliation"
	client.TransactionRevenue = "25.00"
	client.TransactionTax = "1.00"
	client.CurrencyCode = "TRY"

	api.Send(client)
}
