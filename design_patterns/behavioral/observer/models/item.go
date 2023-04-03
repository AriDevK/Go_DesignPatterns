package models

import "fmt"

type Item struct {
	observers []IObserver
	name      string
	available bool
}

func NewItem(name string) *Item {
	return &Item{name: name}
}

func (i *Item) UpdateAvailable() {
	fmt.Printf("The item %s is available\n", i.name)
	i.available = true
	i.Broadcast()
}

func (i *Item) Register(observer IObserver) {
	i.observers = append(i.observers, observer)
}

func (i *Item) Broadcast() {
	for _, observer := range i.observers {
		observer.UpdateValues(i.name)
	}
}
