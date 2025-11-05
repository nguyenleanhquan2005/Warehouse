package entity

import "gorm.io/gorm"

type Warehouse struct {
	gorm.Model
	Name string
	Adress string
	WarehouseType string
	
}
