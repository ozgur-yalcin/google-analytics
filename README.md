[![Linux Build Status](https://travis-ci.org/OzqurYalcin/google-analytics.svg?branch=master)](https://travis-ci.org/OzqurYalcin/google-analytics) [![Windows Build Status](https://ci.appveyor.com/api/projects/status/q7ugwfufg8o55fj4?svg=true)](https://ci.appveyor.com/project/OzqurYalcin/google-analytics) [![Build Status](https://circleci.com/gh/OzqurYalcin/google-analytics.svg?style=svg)](https://circleci.com/gh/OzqurYalcin/google-analytics)

# Google-Analytics
An easy-to-use Google Analytics API (v1) via Measurement Protocol with golang

# Security
If you discover any security related issues, please email ozguryalcin@outlook.com instead of using the issue tracker.

# License
The MIT License (MIT). Please see License File for more information.

# Installation
```bash
go get github.com/OzqurYalcin/google-analytics
```

# Measuring Purchases Example
```go
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
```
