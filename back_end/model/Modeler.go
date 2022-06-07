package model

// Model 通用方法
type Modeler interface {
	Create() (ok bool)
	Update() (ok bool)
	Delete() (ok bool)
}
