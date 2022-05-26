package model

import "gorm.io/gorm"

// Model 通用方法
type Modeler interface {
	Create() (tx *gorm.DB, ok bool)
	Update() (tx *gorm.DB, ok bool)
	Delete() (tx *gorm.DB, ok bool)
}
