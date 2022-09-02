package main

import(
	"log"
	"order/order"
	"order/config"
)

func main() {
	config.ConfigMode()
	options := config.Reader(config.GetMode())
	log.Printf("option : %v", options)
	order.OrderDBInit()
	go order.InitProducer(config.GetMode())
	e := order.RouteInit()

	e.Logger.Fatal(e.Start(":8081"))
}
