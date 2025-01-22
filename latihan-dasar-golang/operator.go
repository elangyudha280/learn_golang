package main

import "fmt"

func main() {

	//! OPERATOR ARITMATIK
	var nilai1 int = 10
	var nilai2 int = 5

	// penjumlahan
	fmt.Println("penjumlahan:")
	// pengurangan
	fmt.Println(nilai1 + nilai2)
	// pengurangan
	fmt.Println(nilai1 - nilai2)
	// perkalian
	fmt.Println(nilai1 * nilai2)
	// pembagian
	fmt.Println(nilai1 / nilai2)
	// modulus
	fmt.Println(nilai1 % nilai2)

	//! OPERATOR ASSINGMENT
	var nilai3 int = 10

	// tambah
	nilai3 += 5
	// kurang
	nilai3 -= 3
	// kali
	nilai3 *= 2
	// bagi
	nilai3 /= 3
	// modulus
	nilai3 %= 2

	fmt.Println("nilai3 setelah melakukan serangkai opeartor assignment:", nilai3)

	//! OPERATOR PERBANDINGAN
	var nilai4 int = 10
	var nilai5 int = 5

	// sama dengan
	fmt.Println("sama dengan:", nilai4 == nilai5)
	// tidak sama dengan
	fmt.Println("tidak sama dengan:", nilai4 != nilai5)
	// lebih kecil
	fmt.Println("lebih kecil:", nilai4 < nilai5)
	// lebih besar
	fmt.Println("lebih besar:", nilai4 > nilai5)
	// lebih kecil atau sama
	fmt.Println("lebih kecil atau sama:", nilai4 <= nilai5)
	// lebih besar atau sama
	fmt.Println("lebih besar atau sama:", nilai4 >= nilai5)

	//! OPERATOR LOGIKA
	var nilai6 int = 10
	var nilai7 int = 5

	// AND
	fmt.Println("AND:", nilai6 > 5 && nilai7 < 10)
	// OR
	fmt.Println("OR:", nilai6 > 5 || nilai7 < 10)
	// NOT
	fmt.Println("NOT:", !true)
}
