package query_builder

import "gorm.io/gorm"

type I interface {
	Query(db *gorm.DB) *gorm.DB
}
