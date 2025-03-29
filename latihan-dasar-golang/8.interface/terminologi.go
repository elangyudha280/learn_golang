package main

import "fmt"

//! Implementasi Interface dengan Struct
// Definisi interface
type Speaker interface {
	Speak() string
}

// Struct yang mengimplementasikan interface
//TODO: NOTES INTERFACE INI SUDAH DI ANGGGAP MENERAPKAN INTERFACE Speaker karena di struct object ini memiliki method yg bernama Speak() string
type Person struct {
	Name string
}

func (p Person) Speak() string {
	return "Halo, saya " + p.Name
}

//! Implementasi Interface dengan Pointer Receiver
type Resettable interface {
	Reset()
}
type Counter struct {
	Value int
}

// Menggunakan pointer agar nilai bisa diubah
func (c *Counter) Reset() {
	c.Value = 0
}

func main() {
	// use interface object struct in variable
	var human1 Speaker = Person{
		Name: "raka",
	}
	fmt.Println(human1.Speak())

	// use pointer receiver in struct
	// c := Counter{Value: 10}
	// fmt.Println("Sebelum reset:", c.Value)

	// var r Resettable = &c // Harus pakai & karena method menerima pointer
	// r.Reset()

	// // fmt.Println("Setelah reset:", c.Value)

	//! variable dengan interface kosong atau any
	// var global_data interface{} //-> ini artinya varibable ini bisa bernilai dengan tipe data apapun
	// ||
	var global_data any
	global_data = 10
	fmt.Printf("Type: %T, Value: %v\n", global_data, global_data)
	global_data = "hellow world"
	fmt.Printf("Type: %T, Value: %v\n", global_data, global_data)
}
