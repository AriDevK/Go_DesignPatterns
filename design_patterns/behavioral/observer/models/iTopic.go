package models

type ITopic interface {
	Register(IObserver)
	Broadcast()
}
