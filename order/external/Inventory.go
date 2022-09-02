package external

type Inventory struct {
	Id int `gorm:"primaryKey" json:"id" type:"int"`
	Stock int `json:"stock" type:"int"`
}

func (self *Inventory) getPrimaryKey() int {
	// FIXME if PrimaryKey is multi value, than change this method
	return self.Id
}
