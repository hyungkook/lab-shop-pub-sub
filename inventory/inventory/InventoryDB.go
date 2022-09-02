package inventory
import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type InventoryDB struct{
	db *gorm.DB
}

var inventoryrepository *InventoryDB

func InventoryDBInit() {
	var err error
	inventoryrepository = &InventoryDB{}
	inventoryrepository.db, err = gorm.Open(sqlite.Open("Inventory_table.db"), &gorm.Config{})
	
	if err != nil {
		panic("DB Connection Error")
	}
	inventoryrepository.db.AutoMigrate(&Inventory{})

}

func InventoryRepository() *InventoryDB {
	return inventoryrepository
}

func (self *InventoryDB)save(entity interface{}) error {
	
	tx := self.db.Create(entity)

	if tx.Error != nil {
		log.Print(tx.Error)
		return tx.Error
	}
	return nil
}

func (self *InventoryDB)GetList() []Inventory{
	
	entities := []Inventory{}
	self.db.Find(&entities)

	return entities
}

func (self *InventoryDB)GetID(id int) *Inventory{
	entity := &Inventory{}
	self.db.Where("id = ?", id).First(entity)

	return entity
}

func (self *InventoryDB) Delete(entity *Inventory) error{
	err2 := self.db.Delete(&entity).Error
	return err2
}

func (self *InventoryDB) Update(id int, params map[string]string) error{
	entity := &Inventory{}
	err1 := self.db.Where("id = ?", id).First(entity).Error
	if err1 != nil {
		return err1
	}else {
		update := &Inventory{}
		ObjectMapping(update, params)

		err2 := self.db.Model(&entity).Updates(update).Error
		return err2
	}

}