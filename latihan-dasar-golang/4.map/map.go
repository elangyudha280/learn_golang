package main

import "fmt"

func main() {

	//todo: deklarasi map langsung
	person := map[string]string{
		"name": "person1",
		"umur": "19",
	}

	fmt.Println(person)

	//! cek key di map
	name2, ifexists := person["name"]
	fmt.Println(name2, ifexists)
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

	//! buat map dengan function make || manual
	person2 := make(map[string]string)
	person2["name"] = "person2"

	fmt.Println(person2)
}
