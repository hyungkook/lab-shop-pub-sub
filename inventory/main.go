package main

import(
	"log"
	"inventory/inventory"
	"inventory/config"
)

func main() {
	config.ConfigMode()
	options := config.Reader(config.GetMode())
	log.Printf("option : %v", options)
	inventory.InventoryDBInit()
	go inventory.InitProducer(config.GetMode())
	// view 와 같이 사용시 InitConsumer 가 중복으로 호출될수 있으니, 하나는 삭제 필요
	go inventory.InitConsumer(config.GetMode())
	e := inventory.RouteInit()

	e.Logger.Fatal(e.Start(":8082"))
}
