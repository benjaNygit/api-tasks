package models

type Model interface {
	Create() bool
	Read() Model
	Update() Model
	Delete() bool
}
