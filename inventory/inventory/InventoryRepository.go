package inventory

import (
	"net/http"
	"strconv"
	"github.com/labstack/echo"
)

func (self *Inventory) Get(c echo.Context) error {
	repository := InventoryRepository()
	entities := repository.GetList()
	return c.JSON(http.StatusOK, entities)
}

func (self *Inventory) GetbyId(c echo.Context) error{
	repository := InventoryRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	self = repository.GetID(id)

	return c.JSON(http.StatusOK, self)
}

func (self *Inventory) Persist(c echo.Context) error{
	repository := InventoryRepository()
	params := make(map[string] string)
	
	c.Bind(&params)
	ObjectMapping(self, params)

	err := repository.save(self)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}else{
		return c.JSON(http.StatusOK, self)
	}
}

func (self *Inventory) Put(c echo.Context) error{
	repository := InventoryRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	params := make(map[string] string)
	
	c.Bind(&params)

	err := repository.Update(id, params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	} else {
		entity := repository.GetID(id)
		return c.JSON(http.StatusOK, entity)
	}
}

func (self *Inventory) Remove(c echo.Context) error{
	repository := InventoryRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	self = repository.GetID(id)

	err := repository.Delete(self)

	return c.JSON(http.StatusOK, err)
}