package main

import "fmt"

//! basic function
func sayHello() {
	fmt.Println("Hello, World!")
}

//! function with parameter
func Greeting(name string, umur string) {
	fmt.Println("Hello, saya ", name, " umur saya ", umur)
}

//! function with return value
func sum(value1, value2 int) int {
	hasil := value1 + value2
	return hasil
}

//! function with multiple return value
func calculate(value1, value2 int) (int, int) {
	return value1 + value2, value1 - value2
}

func main() {
	// sayHello()

	// Greeting("user", "18") // output: Hello, saya  user  umur saya  18

	fmt.Println(sum(5, 10)) // output: 15

	// hasil penjumlahan, pengurangan
	_, perkalian := calculate(5, 10)
	print(perkalian)
}
