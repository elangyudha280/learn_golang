package main

import "fmt"

func main() {

	//todo: deklarasi map manual
	person := map[string]string{
		"name": "person1",
		"umur": "19",
	}

	fmt.Println(person)

	//! akses data map
	// fmt.Println(person["name"])

	//!  panjang data map
	// fmt.Println(len(person))

	//!  ubah data map
	// person["name"] = "person2"
	// fmt.Println(person)

	//! delete key pada map
	delete(person, "umur")
	fmt.Println(person)

	//! buat map dengan function make
	person2 := make(map[string]string)
	person2["name"] = "person2"

	fmt.Println(person2)
}
