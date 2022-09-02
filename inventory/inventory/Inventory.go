package inventory

type Inventory struct {
	Id int `gorm:"primaryKey" json:"id" type:"int"`
	Stock int `json:"stock" type:"int"`
}

func (self *Inventory) AfterCreate(tx *gorm.DB) (err error){
	return nil
}
