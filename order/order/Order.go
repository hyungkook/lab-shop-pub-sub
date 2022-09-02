package order

import (
	"gopkg.in/jeevatkm/go-model.v1"
	
	"gorm.io/gorm"
)
type Order struct {
	Id int `gorm:"primaryKey" json:"id" type:"int"`
	ProductId string `json:"productId" type:"string"`
	Qty int `json:"qty" type:"int"`
	CustomerId string `json:"customerId" type:"string"`
	Amount int `json:"amount" type:"int"`
}

func (self *Order) AfterCreate(tx *gorm.DB) (err error){
	orderPlaced := NewOrderPlaced()
	model.Copy(orderPlaced, self)

	Publish(orderPlaced)
	
	
	return nil
}
func (self *Order) BeforeCreate(tx *gorm.DB) (err error){
	return nil
}
