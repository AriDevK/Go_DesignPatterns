package main

import m "Go_DesignPatterns_Platzi/design_patterns/behavioral/observer/models"

func main() {
	nvidiaItem := m.NewItem("RTX 3080")
	firstObserver := m.NewEmailClient()
	secondObserver := m.NewEmailClient()

	nvidiaItem.Register(firstObserver)
	nvidiaItem.Register(secondObserver)
	nvidiaItem.UpdateAvailable()
}
