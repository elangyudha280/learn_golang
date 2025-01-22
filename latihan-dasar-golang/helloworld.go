// package main

// import "fmt"

// func main() {
// 	// fmt.Println("hello world golang")
// 	// fmt.Print("hello world golang")
// 	// fmt.Printf("hello world golang")
// 	var hello string = "hello"

// 	fmt.Println(hello, len(hello), hello[1])

// 	//* variable var explicit tipe data
// 	var angka int = 1 //-> variablnya bisa di ubah atau deklarasi ulang dan ini bisa juga untuk membuat variabel global
// 	//* variable var implicit tipe data
// 	var a = "user1" //-> ini artinya nilainya akan bertipe data sesuai yg kita tentukan
// 	//* variable var tidak memiliki data haya deklarasi
// 	var dekVar bool //-> ini artinya nilai defaultnya sesuai nilai default dari tipe datanya

// 	angka = 2

// 	//* variable dengan const explicit tipe data
// 	const imutableAngka int = 4 //-> variablnya TIDAK bisa di ubah atau deklarasi ulang
// 	// imutableAngka = 3 //-> ini tidak bisa karena menggunakan keyword const
// 	//* variable var implicit tipe data
// 	const imutableImplAngka = 1 //-> ini artinya nilainya akan bertipe data sesuai yg kita tentukan

// 	// VARIABLE DENGAN SHORT DECLARATION
// 	varShorDec := 123

// 	//* multiple variable with var
// 	var (
// 		var1 int = 1
// 		var2     = "2"
// 		var3 bool
// 	)
// 	var2 = "3"

// 	//* multiple variable with const
// 	// const (
// 	// 	const1 int = 1
// 	// 	const2     = "2"
// 	// 	const3 bool
// 	// )

// 	//! type declaration
// 	type nama_user = string

// 	var user1 nama_user = "user1"

// 	fmt.Println(user1, a, dekVar)
// 	fmt.Println(angka, a, dekVar)
// 	fmt.Println(var1, var2, var3)
// 	fmt.Println(imutableAngka)
// 	fmt.Println("variable yg dibuad dengan short deklarasi", varShorDec)
// }
