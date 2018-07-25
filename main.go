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
	client.UserID = ""

	client.CurrencyCode = "TRY"
	response := api.Send(client)
	fmt.Println(response)
}
