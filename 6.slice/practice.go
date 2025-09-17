package main

import "fmt"

func main() {
	//! buah
	// var buahSlice = []string{"mangga", "apel", "jeruk", "semangka", "melon", "anggur", "kiwi"}
	// println("data buah awal", buahSlice[0])
	// println("data  buah akhir", buahSlice[len(buahSlice)-1])
	// //! panjang data buah
	// println("panjang data buah", len(buahSlice))
	// //! kapasitas data buah
	// println("kapasitas data buah", cap(buahSlice))

	// var angka = []int{1, 2, 3, 4, 5}
	// angka2 := append(angka, 6, 7, 8, 9, 10)
	// //! panjang data angka
	// println("panjang data angka", len(angka))
	// //! kapasitas data angka
	// println("kapasitas data angka", cap(angka))

	// println("panjang data angka2", len(angka2))
	// //! kapasitas data angka
	// println("kapasitas data angka2", cap(angka2))

	var angka = make([]int, 2, 5)
	 angka = append(angka, 1, 2, 3,4,5,6)
	fmt.Println("lengt data", len(angka))
	fmt.Println("capacity data angka ", cap(angka))
}