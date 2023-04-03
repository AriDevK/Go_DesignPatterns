package main

func main() {
	laptop, _ := GetComputerFactory("Laptop")
	desktop, _ := GetComputerFactory("Desktop")

	PrintProductNameAndStock(laptop)
	PrintProductNameAndStock(desktop)
}
