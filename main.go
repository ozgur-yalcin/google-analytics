package main

import (
	"net/http"
	"time"

	"github.com/OzqurYalcin/google-analytics/src"
	"github.com/google/uuid"
)

func main() {
	http.HandleFunc("/", view)
	server := http.Server{Addr: ":8080", ReadTimeout: 30 * time.Second, WriteTimeout: 30 * time.Second}
	server.ListenAndServe()
}

func view(w http.ResponseWriter, r *http.Request) {
	api := new(ga.API)
	api.Lock()
	defer api.Unlock()
	api.UserAgent = r.UserAgent()
	api.ContentType = "application/x-www-form-urlencoded"

	client := new(ga.Client)
	client.ProtocolVersion = "1"
	client.ClientID = uuid.New().String()
	client.TrackingID = "UA-xxxxxxxx-x"
	client.HitType = "pageview"
	client.DocumentLocationURL = "https://www.example.com/payment"
	client.DocumentTitle = "Payment"
	client.DocumentEncoding = "UTF-8"

	product := new(ga.Product)
	product.SKU = "P1234"
	product.Name = "product name"
	product.Brand = "product brand"
	product.Price = "1.00"
	product.Quantity = "1"
	client.Products = append(client.Products, product)
	client.ProductAction = "purchase"
	client.TransactionID = "T1234"
	client.TransactionRevenue = "1.00"
	client.CurrencyCode = "TRY"

	if r.URL.Path == "/" {
		api.Send(client)
	}
}
