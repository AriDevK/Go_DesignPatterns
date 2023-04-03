package models

type IObserver interface {
	GetId() string
	UpdateValues(string)
}
