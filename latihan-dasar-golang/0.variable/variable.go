package main

import "fmt"

//! 1. variable adlaah sebuah tempat untuk menyimpan sebuah nilai sehingga ketika kita ingin menggunakan nilai tersebut kita tinggal panggil variablenya

func main() {
	//!!! membuat variable degan keyword var
	//!!  deklarasi tanpa assigment
	var angka1 int
	angka1 = 100
	angka1 += 100

	//!! deklarasi langsung assigment
	var angka2 int = 100

	//!!! deklarasi variable dengan short deklarasi
	angka3 := angka1+angka2

	//!! multiple deklarasi
	var (
		nama = "user1"
		nama2 = "user2"
		nama3 = "user3"
	)  

	fmt.Println(angka3)
	fmt.Println(nama,nama2,nama3)


	// 
}