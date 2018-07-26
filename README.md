# Google-Analytics

# Security
If you discover any security related issues, please email ozguryalcin@outlook.com instead of using the issue tracker.

# License
The MIT License (MIT). Please see License File for more information.

# Installation
```bash
go get github.com/OzqurYalcin/google-analytics
```

# Usage
```go
package main

import (
	"fmt"

	"github.com/OzqurYalcin/google-analytics/src"
)

func main() {
	api := new(ga.API)
	api.Lock()
	defer api.Unlock()
	client := new(ga.Client)
	client.ProtocolVersion = "1"
	client.TrackingID = "UA-xxxxxxxx-xx"
	client.HitType = "transaction"

	product := new(ga.Product)
	product.Name = "-"
	product.Price = "0.00"
	product.Quantity = "1"
	product.Brand = "-"
	product.Action = "purchase"
	client.Products = append(client.Products, product)

	client.TransactionID = ""
	client.TransactionRevenue = "0.00"

	client.CurrencyCode = "USD"
	response := api.Send(client)
	fmt.Println(response)
}
```
